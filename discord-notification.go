package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"os"
)

func SendToDiscord(message string) {
	if !*EnableDiscordNotifications {
		return
	}

	url := os.Getenv("DISCORD_WEBHOOK_URL")
	if url == "" {
		log.Println("Failed to send discord message: url not defined")
		return
	}

	body, err := json.Marshal(map[string]string{
		"content":  message,
		"username": "dioggy bot superviser",
	})
	if err != nil {
		log.Printf("Failed to send discord message: %s\n", err.Error())
		return
	}

	_, err = http.Post(url, "application/json", bytes.NewBuffer(body))
	if err != nil {
		log.Printf("Failed to send discord message: %s\n", err.Error())
		return
	}
}
