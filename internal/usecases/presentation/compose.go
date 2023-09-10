package presentation

import (
	"chess-daily-puzzle/internal/models"
	"encoding/json"
	"fmt"
)

const endpointPuzzleTraining = "https://lichess.org/training/"

func ComposePayload(gameID, gamePicURL string) map[string]interface{} {
	if gameID == "" || gamePicURL == "" {
		return nil
	}

	gameURL := endpointPuzzleTraining + gameID

	pl := models.Payload{
		Username: Username,
		Text:     fmt.Sprintf(Message, gameURL),
		IconURL:  IconURL,
		Attachments: []map[string]interface{}{
			{
				"image_url": gamePicURL,
			},
		},
	}

	plEncoded, err := json.Marshal(pl)
	if err != nil {
		return nil
	}

	var result map[string]interface{}
	err = json.Unmarshal(plEncoded, &result)
	if err != nil {
		return nil
	}

	return result
}
