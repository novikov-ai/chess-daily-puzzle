package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
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

	picURL, err := pgn.GetPictureURL(puzzle.Puzzle.Fen)
	if err != nil {
		telegram.LogError(tgLogger, "Can't get picture from FEN", err)
		log.Fatal("Can't get picture from FEN:", err)
	}

	log.Printf("Got puzzle %s, picture: %s", puzzle.Puzzle.ID, picURL)

	payload := presentation.ComposePayload(puzzle.Puzzle.ID, picURL)
	if payload == nil {
		err = errors.New("empty puzzle id or picture url")
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

	if resp.StatusCode < http.StatusOK || resp.StatusCode >= http.StatusMultipleChoices {
		body, _ := io.ReadAll(io.LimitReader(resp.Body, 200))
		err = fmt.Errorf("unexpected status code: %d, body: %q", resp.StatusCode, body)
		telegram.LogError(tgLogger, "Webhook rejected the puzzle", err)
		log.Fatal("Webhook rejected the puzzle:", err)
	}

	msg := fmt.Sprintf("Puzzle %s was sent successfully!", puzzle.Puzzle.ID)
	log.Println(msg)
	telegram.LogInfo(tgLogger, msg)
}
