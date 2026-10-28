package pgn

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetPictureURL(t *testing.T) {
	testCases := []struct {
		name     string
		fen      string
		wantErr  bool
		expected string
	}{
		{
			name:     "Black moves",
			fen:      "r7/pp1b4/2Pkpp1Q/2q3r1/8/3B1p1P/PP3PP1/R2R2K1 b - - 0 1",
			expected: "https://lichess1.org/export/fen.gif?color=black&fen=r7%2Fpp1b4%2F2Pkpp1Q%2F2q3r1%2F8%2F3B1p1P%2FPP3PP1%2FR2R2K1+b+-+-+0+1&piece=cburnett&theme=brown&variant=standard",
		},
		{
			name:     "White moves",
			fen:      "4K3/3P4/pb6/2p3p1/2P3k1/1P6/P7/8 w - - 1 1",
			expected: "https://lichess1.org/export/fen.gif?color=white&fen=4K3%2F3P4%2Fpb6%2F2p3p1%2F2P3k1%2F1P6%2FP7%2F8+w+-+-+1+1&piece=cburnett&theme=brown&variant=standard",
		},
		{name: "empty", fen: "", wantErr: true},
		{name: "unknown side", fen: "8/8/8/8/8/8/8/8 x - - 0 1", wantErr: true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result, err := GetPictureURL(tc.fen)

			assert.Equal(t, tc.wantErr, err != nil)
			assert.Equal(t, tc.expected, result)
		})
	}
}
