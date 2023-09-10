package presentation

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestComposePayload(t *testing.T) {
	testCases := []struct {
		name       string
		gameID     string
		gamePicURL string
		expected   map[string]interface{}
	}{
		{
			name:       "ok",
			gameID:     "12345",
			gamePicURL: "https://example.com/game.png",
			expected: map[string]interface{}{
				"username": "Daily Puzzle",
				"text":     "[Найдите](https://lichess.org/training/12345) лучшее продолжение!",
				"icon_url": "https://lichess1.org/assets/_44IzGj/logo/lichess-favicon-128.png",
				"attachments": []interface{}{
					map[string]interface{}{
						"image_url": "https://example.com/game.png",
					},
				},
			},
		},
		{
			name:       "Empty URL",
			gameID:     "",
			gamePicURL: "https://example.com/game.png",
			expected:   nil,
		},
		{
			name:       "Empty game pic URL",
			gameID:     "12345",
			gamePicURL: "",
			expected:   nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := ComposePayload(tc.gameID, tc.gamePicURL)
			assert.Equal(t, result, tc.expected)
		})
	}
}
