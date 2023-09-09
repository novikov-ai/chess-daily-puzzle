package presentation

import (
	"chess-daily-puzzle/internal/models"
	"encoding/json"
	"fmt"
	"log"
)

const endpointPuzzleTraining = "https://lichess.org/training/"

func ComposePayload(gameURL, gamePicURL string) map[string]interface{} {
	if gameURL == "" || gamePicURL == "" {
		return nil
	}

	pl := models.Payload{
		Username: "Daily Puzzle",
		Text:     fmt.Sprintf("[Найдите](%s%s) лучшее продолжение!", endpointPuzzleTraining, gameURL),
		IconURL:  "https://lichess1.org/assets/_44IzGj/logo/lichess-favicon-128.png",
		Attachments: []map[string]interface{}{
			{
				"image_url": gamePicURL,
			},
		},
	}

	plEncoded, err := json.Marshal(pl)
	if err != nil {
		log.Fatalf("Can't compose a payload: %s", err)
	}

	var result map[string]interface{}
	err = json.Unmarshal(plEncoded, &result)
	if err != nil {
		log.Fatalf("Can't compose a payload: %s", err)
	}

	return result
}
