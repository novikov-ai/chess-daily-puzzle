package presentation

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestComposePayload(t *testing.T) {
	testCases := []struct {
		name       string
		gameURL    string
		gamePicURL string
		expected   map[string]interface{}
	}{
		{
			name:       "ok",
			gameURL:    "https://example.com/game",
			gamePicURL: "https://example.com/game.png",
			expected: map[string]interface{}{
				"username": "Coach",
				"text":     "[Найдите](https://example.com/game) лучшее продолжение!",
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
			gameURL:    "",
			gamePicURL: "https://example.com/game.png",
			expected:   nil,
		},
		{
			name:       "Empty game pic URL",
			gameURL:    "https://example.com/game",
			gamePicURL: "",
			expected:   nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := ComposePayload(tc.gameURL, tc.gamePicURL)
			assert.Equal(t, result, tc.expected)
		})
	}
}
