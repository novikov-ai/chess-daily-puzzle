package fetching

import (
	"encoding/json"
	"io"
	"net/http"

	"chess-daily-puzzle/internal/models"
)

const apiPuzzleDaily = "https://lichess.org/api/puzzle/daily"

func DailyPuzzle() (models.Puzzle, error) {
	resp, err := http.Get(apiPuzzleDaily)
	if err != nil {
		return models.Puzzle{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return models.Puzzle{}, err
	}

	var puzzle models.Puzzle
	err = json.Unmarshal(body, &puzzle)
	if err != nil {
		return models.Puzzle{}, err
	}

	return puzzle, nil
}
