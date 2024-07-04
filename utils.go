// utils.go
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const (
	apiBaseURL = "https://api.avito.ru"
	authToken  = "Bearer access_token"
)

func sendMessage(userID int, chatID, messageText string) error {
	url := fmt.Sprintf("%s/messenger/v1/accounts/%d/chats/%s/messages", apiBaseURL, userID, chatID)

	sendMessageRequest := SendMessageRequest{
		Type: "text",
	}
	sendMessageRequest.Message.Text = messageText

	requestBody, err := json.Marshal(sendMessageRequest)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", authToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("failed to send message: %s", string(body))
	}

	return nil
}

func subscribeWebhook() error {
	url := fmt.Sprintf("%s/messenger/v3/webhook", apiBaseURL)

	requestBody, err := json.Marshal(map[string]string{
		"url": "https://your-server.com/webhook",
	})
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", authToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("failed to subscribe webhook: %s", string(body))
	}

	return nil
}
