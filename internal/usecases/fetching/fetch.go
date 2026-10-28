package fetching

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"chess-daily-puzzle/internal/models"
	"chess-daily-puzzle/internal/usecases/retry"
)

const apiPuzzleDaily = "https://lichess.org/api/puzzle/daily"

// Предуловия: нет
// Постусловия: получена моделька пазла или ошибка, если были проблемы с формированием пазла
func DailyPuzzle() (models.Puzzle, error) {
	return retry.Do(retry.DefaultConfig(), func() (models.Puzzle, error) {
		return fetchPuzzle()
	})
}

func fetchPuzzle() (models.Puzzle, error) {
	resp, err := http.Get(apiPuzzleDaily)
	if err != nil {
		return models.Puzzle{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return models.Puzzle{}, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

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
