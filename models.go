// models.go
package main

type WebhookMessage struct {
    AuthorID int    `json:"author_id"`
    ChatID   string `json:"chat_id"`
    Content  struct {
        Text string `json:"text"`
    } `json:"content"`
    Created int    `json:"created"`
    ID      string `json:"id"`
    Type    string `json:"type"`
    UserID  int    `json:"user_id"`
}

type SendMessageRequest struct {
    Message struct {
        Text string `json:"text"`
    } `json:"message"`
    Type string `json:"type"`
}