package presentation

import (
	"chess-daily-puzzle/internal/models"
	"encoding/json"
	"fmt"
	"log"
)

func ComposePayload(gameURL, gamePic string) map[string]interface{} {
	if gameURL == "" || gamePic == "" {
		return nil
	}

	pl := models.Payload{
		Username: "Coach",
		Text:     fmt.Sprintf("[Найдите](%s) лучшее продолжение!", gameURL),
		IconURL:  "https://lichess1.org/assets/_44IzGj/logo/lichess-favicon-128.png",
		Attachments: []map[string]interface{}{
			{
				"image_url": gamePic,
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
