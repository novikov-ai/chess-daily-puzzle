package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"log"
	"net/http"
	"os"

	"chess-daily-puzzle/internal/usecases/telegram"

	"chess-daily-puzzle/internal/usecases/fetching"
	"chess-daily-puzzle/internal/usecases/pgn"
	"chess-daily-puzzle/internal/usecases/presentation"

	"github.com/joho/godotenv"
)

var envFlag string

func main() {
	flag.StringVar(&envFlag, "e", "debug", "environment (prod/dev)")
	flag.Parse()

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Can't find the config file")
	}

	tgLogger, err := telegram.NewLogger()
	if err != nil {
		log.Println("Can't create telegram logger:", err)
	}

	webhookURL := ""
	switch envFlag {
	case "prod":
		webhookURL = os.Getenv("MATTERMOST_WEBHOOK_URL")
	case "debug":
		webhookURL = os.Getenv("MATTERMOST_WEBHOOK_URL_DEBUG")
	}

	if webhookURL == "" {
		telegram.LogError(tgLogger, "Webhook url is empty", nil)
		log.Fatal("Webhook url is empty")
	}

	log.Println("Start fetching a new daily puzzle...")

	puzzle, err := fetching.DailyPuzzle()
	if err != nil {
		telegram.LogError(tgLogger, "Can't fetch daily puzzle", err)
		log.Fatal("Can't fetch daily puzzle:", err)
	}

	picURL, err := pgn.GetPictureURL(puzzle.Game.Pgn)
	if err != nil {
		telegram.LogError(tgLogger, "Can't get picture from PGN", err)
		log.Fatal("Can't get picture from PGN:", err)
	}

	payload := presentation.ComposePayload(puzzle.Puzzle.ID, picURL)
	if payload == nil {
		telegram.LogError(tgLogger, "Error composing payload", err)
		log.Fatal("Error composing payload:", err)
	}

	payloadJSON, err := json.Marshal(payload)
	if err != nil {
		telegram.LogError(tgLogger, "Error creating JSON payload", err)
		log.Fatal("Error creating JSON payload:", err)
	}

	resp, err := http.Post(webhookURL, "application/json", bytes.NewBuffer(payloadJSON))
	if err != nil {
		telegram.LogError(tgLogger, "Error sending webhook request", err)
		log.Fatal("Error sending webhook request:", err)
	}
	defer resp.Body.Close()

	log.Println("Puzzle was sent successfully!")
	tgLogger.Info("Puzzle was sent successfully!")
}
