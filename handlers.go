// handlers.go
package main

import (
	"encoding/json"
	"io"
	"net/http"
)

func handleWebhook(w http.ResponseWriter, r *http.Request) {
	var webhookMessage WebhookMessage
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Unable to read request body", http.StatusBadRequest)
		return
	}

	err = json.Unmarshal(body, &webhookMessage)
	if err != nil {
		http.Error(w, "Unable to parse JSON", http.StatusBadRequest)
		return
	}

	err = sendMessage(webhookMessage.UserID, webhookMessage.ChatID, "100-7!")
	if err != nil {
		http.Error(w, "Unable to send message", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
