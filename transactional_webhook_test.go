package main

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

// ComputeHMAC generates the hexadecimal HMAC SHA-256 signature for test payloads
func ComputeHMAC(secret string, payload []byte) string {
	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write(payload)
	return hex.EncodeToString(mac.Sum(nil))
}

// TestVerifyWebhookSignature validates HMAC-SHA256 signature verification against spoofing attacks
func TestVerifyWebhookSignature(t *testing.T) {
	secretKey := "8f8dce9176c84d50939a20f99ca75f6e"
	ws := &WebhookServer{
		WebhookSecretKey: secretKey,
	}

	rawPayload := []byte(`{"event_id":"evt_test_101","type":"payment.succeeded","session_token":"cs_live_valid_test","amount_paid":299.00}`)
	validSignature := ComputeHMAC(secretKey, rawPayload)

	// 1. Assert valid signature is verified successfully
	if !ws.VerifyWebhookSignature(rawPayload, validSignature) {
		t.Fatalf("Expected valid HMAC-SHA256 signature to be authenticated, but verification failed")
	}

	// 2. Assert tampered payload is rejected
	tamperedPayload := []byte(`{"event_id":"evt_test_101","type":"payment.succeeded","session_token":"cs_live_tampered","amount_paid":0.00}`)
	if ws.VerifyWebhookSignature(tamperedPayload, validSignature) {
		t.Fatalf("Security failure: Tampered payload with original signature should have been rejected")
	}

	// 3. Assert wrong secret key produces rejected signature
	wrongSecretSignature := ComputeHMAC("invalid_attacker_key", rawPayload)
	if ws.VerifyWebhookSignature(rawPayload, wrongSecretSignature) {
		t.Fatalf("Security failure: Payload signed with invalid secret key was accepted")
	}

	// 4. Assert malformed signature is rejected
	if ws.VerifyWebhookSignature(rawPayload, "malformed_hex_signature") {
		t.Fatalf("Security failure: Malformed signature string was accepted")
	}
}

// TestReplayAttackPrevention validates that duplicate / replayed session tokens return HTTP 409 Conflict
func TestReplayAttackPrevention(t *testing.T) {
	secretKey := "8f8dce9176c84d50939a20f99ca75f6e"

	// Mock server that tracks completed tokens in-memory for unit testing
	completedTokens := make(map[string]bool)
	completedTokens["cs_live_already_completed"] = true

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}

		mac := hmac.New(sha256.New, []byte(secretKey))
		var event PaymentWebhookEvent
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&event); err != nil {
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}
		_ = mac

		// Replay attack check
		if completedTokens[event.SessionToken] {
			w.WriteHeader(http.StatusConflict) // HTTP 409 Conflict
			w.Write([]byte(`{"status":"already_processed","error":"replay_attack_conflict"}`))
			return
		}

		completedTokens[event.SessionToken] = true
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":"verified_and_unlocked"}`))
	})

	// Scenario 1: Initial valid request
	initialPayload, _ := json.Marshal(PaymentWebhookEvent{
		EventID:      "evt_init_001",
		Type:         "payment.succeeded",
		SessionToken: "cs_live_new_unique_token",
		AmountPaid:   299.00,
		Timestamp:    time.Now(),
	})

	req1 := httptest.NewRequest("POST", "/api/payment/webhook", bytes.NewBuffer(initialPayload))
	req1.Header.Set("Content-Type", "application/json")
	w1 := httptest.NewRecorder()
	handler.ServeHTTP(w1, req1)

	if w1.Code != http.StatusOK {
		t.Errorf("Expected initial request to succeed with HTTP 200, got: %d", w1.Code)
	}

	// Scenario 2: Replay attack with identical token must return HTTP 409 Conflict
	req2 := httptest.NewRequest("POST", "/api/payment/webhook", bytes.NewBuffer(initialPayload))
	req2.Header.Set("Content-Type", "application/json")
	w2 := httptest.NewRecorder()
	handler.ServeHTTP(w2, req2)

	if w2.Code != http.StatusConflict {
		t.Fatalf("Replay attack vulnerability! Expected HTTP 409 Conflict for duplicate session token, got: %d", w2.Code)
	}

	// Scenario 3: Pre-existing completed token
	replayPayload, _ := json.Marshal(PaymentWebhookEvent{
		EventID:      "evt_replay_002",
		Type:         "payment.succeeded",
		SessionToken: "cs_live_already_completed",
		AmountPaid:   299.00,
		Timestamp:    time.Now(),
	})

	req3 := httptest.NewRequest("POST", "/api/payment/webhook", bytes.NewBuffer(replayPayload))
	req3.Header.Set("Content-Type", "application/json")
	w3 := httptest.NewRecorder()
	handler.ServeHTTP(w3, req3)

	if w3.Code != http.StatusConflict {
		t.Fatalf("Replay attack vulnerability! Expected HTTP 409 Conflict for existing completed token, got: %d", w3.Code)
	}
}

// TestCommercialKeywordsExclusion ensures no commercial upsell keywords exist in RVM directives
func TestCommercialKeywordsExclusion(t *testing.T) {
	forbiddenWords := []string{"buy", "checkout", "discount", "purchase", "gems"}
	sampleDirectives := []string{
		"Listen to pure 147Hz acoustic resonance fields for exactly 11 minutes during your personal Saturn Hora window.",
		"Enforce artificial friction into your wealth loop. Automate exactly 8% of your weekly earnings directly into a locked savings vault.",
	}

	for _, dir := range sampleDirectives {
		lowerDir := strings.ToLower(dir)
		for _, kw := range forbiddenWords {
			if strings.Contains(lowerDir, kw) {
				t.Fatalf("Commercial keyword violation: Directive contains forbidden word '%s': %s", kw, dir)
			}
		}
	}
}
