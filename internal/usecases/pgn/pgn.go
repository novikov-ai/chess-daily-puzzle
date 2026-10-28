package pgn

import (
	"errors"
	"net/url"
	"strings"
)

const exportFenGif = "https://lichess1.org/export/fen.gif"

// Предуловия: fen — строка в формате Forsyth–Edwards Notation с позицией, с которой начинается пазл
// Постусловия: валидный URL-адрес с корректным расположением доски (с точки зрения стороны, которая ходит) и ошибка либо ее отсутствие
func GetPictureURL(fen string) (string, error) {
	fields := strings.Fields(fen)
	if len(fields) < 2 {
		return "", errors.New("wrong fen format")
	}

	color := ""
	switch fields[1] {
	case "w":
		color = "white"
	case "b":
		color = "black"
	default:
		return "", errors.New("wrong fen format")
	}

	query := url.Values{
		"fen":     []string{fen},
		"color":   []string{color},
		"variant": []string{"standard"},
		"theme":   []string{"brown"},
		"piece":   []string{"cburnett"},
	}

	return exportFenGif + "?" + query.Encode(), nil
}
