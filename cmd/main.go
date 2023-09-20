package main

import (
	"bytes"
	"chess-daily-puzzle/internal/usecases/fetching"
	"chess-daily-puzzle/internal/usecases/pgn"
	"chess-daily-puzzle/internal/usecases/presentation"
	"encoding/json"
	"flag"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

var envFlag string

func main() {
	flag.StringVar(&envFlag, "e", "debug", "environment (prod/dev)")
	flag.Parse()

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Can't find the config file")
	}

	webhookURL := ""
	switch envFlag {
	case "prod":
		webhookURL = os.Getenv("MATTERMOST_WEBHOOK_URL")
	case "debug":
		webhookURL = os.Getenv("MATTERMOST_WEBHOOK_URL_DEBUG")
	}

	if webhookURL == "" {
		log.Fatal("Webhook url is empty")
	}

	log.Println("Start fetching a new daily puzzle...")

	puzzle, err := fetching.DailyPuzzle()
	if err != nil {
		log.Fatal("Can't fetch daily puzzle:", err)
	}

	picURL, err := pgn.GetPictureURL(puzzle.Game.Pgn)
	if err != nil {
		log.Fatal("Can't get picture from PGN:", err)
	}

	payload := presentation.ComposePayload(puzzle.Puzzle.Id, picURL)
	if payload == nil {
		log.Fatal("Error composing payload:", err)
	}

	payloadJSON, err := json.Marshal(payload)
	if err != nil {
		log.Fatal("Error creating JSON payload:", err)
	}

	resp, err := http.Post(webhookURL, "application/json", bytes.NewBuffer(payloadJSON))
	if err != nil {
		log.Fatal("Error sending webhook request:", err)
	}
	defer resp.Body.Close()

	log.Println("Puzzle was sent successfully!")
}
