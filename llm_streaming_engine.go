package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type UserDemographics struct {
	CurrentOccupation  string `json:"current_occupation"`
	GeographicLocation string `json:"geographic_location"`
	SocioEconomicEra   string `json:"socio_economic_era"`
}

type LLMContextPayload struct {
	CalculationMetrics ResolutionPayload `json:"calculation_metrics"`
	RemedialDirectives RVMOutputPayload  `json:"remedial_directives"`
	UserContext        UserDemographics  `json:"user_context"`
}

type StreamOrchestrator struct {
	APIEndpoint string
	APIKey      string
}

func (so *StreamOrchestrator) StreamContextualProse(ctx context.Context, payload LLMContextPayload, w io.Writer) error {
	dataBytes, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("payload generation error: %v", err)
	}

	systemInstruction := `You are the natural language synthesis tier of a high-precision astrology engine.
1. You are strictly forbidden from calculating planetary degrees or modifying numerical data.
2. Apply the ancient structural law of Desha (Location), Kala (Era), and Patra (Individual Context).
3. Translate ancient medieval metaphors into modern terms (e.g., replace 'elephants' with 'venture funding or asset portfolio').
4. Deliver a single, integrated summary. Do not print contradictory statements.`

	requestBody := map[string]interface{}{
		"model": "gemini-2.5-pro",
		"contents": []map[string]interface{}{
			{
				"role": "user",
				"parts": []map[string]interface{}{
					{
						"text": systemInstruction + "\nData: " + string(dataBytes),
					},
				},
			},
		},
	}

	reqJSON, _ := json.Marshal(requestBody)
	req, err := http.NewRequestWithContext(ctx, "POST", so.APIEndpoint, bytes.NewBuffer(reqJSON))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-goog-api-key", so.APIKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	_, err = io.Copy(w, resp.Body)
	return err
}
