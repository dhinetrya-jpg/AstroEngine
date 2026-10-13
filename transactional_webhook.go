package main

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type PaymentWebhookEvent struct {
	EventID      string    `json:"event_id"`
	Type         string    `json:"type"`          // e.g., "payment.succeeded"
	SessionToken string    `json:"session_token"` // Unique lookup token matching database register
	AmountPaid   float64   `json:"amount_paid"`
	Timestamp    time.Time `json:"timestamp"`
}

type WebhookServer struct {
	DB               *sql.DB
	WebhookSecretKey string
}

func (ws *WebhookServer) VerifyWebhookSignature(payload []byte, receivedSignature string) bool {
	mac := hmac.New(sha256.New, []byte(ws.WebhookSecretKey))
	mac.Write(payload)
	expectedSignature := hex.EncodeToString(mac.Sum(nil))
	return hmac.Equal([]byte(expectedSignature), []byte(receivedSignature))
}

func (ws *WebhookServer) HandlePaymentCallback(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	gatewaySignature := r.Header.Get("X-Gateway-Signature")
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Bad Request Ingestion Fault", http.StatusBadRequest)
		return
	}

	// 1. Security Check: Authenticate payload signature against spoofing attempts
	if !ws.VerifyWebhookSignature(body, gatewaySignature) {
		http.Error(w, "Unauthorized Webhook Signature Breach", http.StatusUnauthorized)
		return
	}

	var event PaymentWebhookEvent
	if err := json.Unmarshal(body, &event); err != nil {
		http.Error(w, "Invalid Payload Schema JSON", http.StatusBadRequest)
		return
	}

	// 2. Process Successful Transactions
	if event.Type == "payment.succeeded" {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		tx, err := ws.DB.BeginTx(ctx, nil)
		if err != nil {
			http.Error(w, "Internal Cluster DB Error", http.StatusInternalServerError)
			return
		}

		// Look up the transaction in the database using the unique session token code
		var currentStatus string
		var userID string
		queryLookup := `SELECT status, user_id FROM payment_transactions WHERE session_token = $1 FOR UPDATE`
		err = tx.QueryRowContext(ctx, queryLookup, event.SessionToken).Scan(&currentStatus, &userID)
		if err == sql.ErrNoRows {
			tx.Rollback()
			http.Error(w, "Target Transaction Token Registry Not Found", http.StatusNotFound)
			return
		}

		// Prevent double-processing or replay attack vulnerabilities
		if currentStatus == "COMPLETED" {
			tx.Commit()
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"status":"already_processed"}`))
			return
		}

		// 3. Mutate State Parameters: Move transaction from PENDING to COMPLETED
		queryUpdate := `UPDATE payment_transactions SET status = 'COMPLETED', updated_at = NOW() WHERE session_token = $1`
		_, err = tx.ExecContext(ctx, queryUpdate, event.SessionToken)
		if err != nil {
			tx.Rollback()
			http.Error(w, "Database State Alteration Exception", http.StatusInternalServerError)
			return
		}
		err = tx.Commit()
		if err != nil {
			http.Error(w, "Database Commit Exception", http.StatusInternalServerError)
			return
		}

		// 4. Trigger Real-Time Frontend Unlock Notice
		// In production, this emits an active notification event via Redis Pub/Sub to the SSE loop
		fmt.Printf("[SYSTEM_SUCCESS] Verified token %s. Unlocking Dashboard access profiles for User: %s\n", event.SessionToken, userID)
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status":"verified_and_unlocked"}`))
}

func main() {}
