// main.go
package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/webhook", handleWebhook)

	err := subscribeWebhook()
	if err != nil {
		log.Fatalf("Failed to subscribe webhook: %v", err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Listening on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
