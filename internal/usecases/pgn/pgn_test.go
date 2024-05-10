package pgn

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetPictureURL(t *testing.T) {
	testCases := []struct {
		name     string
		pgn      string
		err      error
		expected string
	}{
		{
			name:     "Black moves",
			pgn:      "e4 e6 d4 d5 Nd2 Nf6 Bd3 Be7 Ngf3 c5 dxc5 Bxc5 O-O O-O Re1 Ng4 Rf1 Nc6 h3 Nf6 Qe2 Re8 Rd1 Qb6 Nb3 dxe4 Nxc5 exf3 Qe3 Nd5 Qe4 Qxc5 Qxh7+ Kf8 Qh8+ Ke7 Qxg7 Bd7 Bg5+ Kd6 c4 Rg8 Qh6 f6 cxd5 Rxg5 dxc6",
			err:      nil,
			expected: "https://lichess1.org/export/fen.gif?fen=r7%2Fpp1b4%2F2Pkpp1Q%2F2q3r1%2F8%2F3B1p1P%2FPP3PP1%2FR2R2K1+b+-+-+0+1&variant=standard&theme=brown&piece=cburnett&color=black",
		},
		{
			name:     "White moves",
			pgn:      "e4 e6 d4 d5 Nd2 Nf6 Bd3 Be7 Ngf3 c5 dxc5 Bxc5 O-O O-O Re1 Ng4 Rf1 Nc6 h3 Nf6 Qe2 Re8 Rd1 Qb6 Nb3 dxe4 Nxc5 exf3 Qe3 Nd5 Qe4 Qxc5 Qxh7+ Kf8 Qh8+ Ke7 Qxg7 Bd7 Bg5+ Kd6 c4 Rg8 Qh6 f6 cxd5 Rxg5 dxc6 Rxg2",
			err:      nil,
			expected: "https://lichess1.org/export/fen.gif?fen=r7%2Fpp1b4%2F2Pkpp1Q%2F2q5%2F8%2F3B1p1P%2FPP3Pr1%2FR2R2K1+w+-+-+0+1&variant=standard&theme=brown&piece=cburnett",
		},
		{
			name:     "error",
			pgn:      "",
			err:      errors.New("error"),
			expected: "",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result, err := GetPictureURL(tc.pgn)

			assert.Equal(t, tc.err == nil, err == nil)
			assert.Equal(t, result, tc.expected)
		})
	}
}
