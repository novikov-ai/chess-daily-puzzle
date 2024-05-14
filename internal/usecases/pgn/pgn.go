package pgn

import (
	"bytes"
	"errors"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

const (
	apiGameImport = "https://lichess.org/api/import"
	exportFenGif  = "lichess1.org/export/fen.gif"
)

const (
	colorWhite = "&color=white"
	colorBlack = "&color=black"
)

func GetPictureURL(pgn string) (string, error) {
	respHTML, err := pgnImportRetrieveHTML(pgn)
	if err != nil {
		return "", err
	}

	picURL := getPositionURL(respHTML)
	if picURL == "" {
		return "", errors.New("picture not found")
	}

	flipBoard, err := blackMove(pgn)
	if err != nil {
		return "", err
	}

	if flipBoard {
		picURL = picUrlWithColorBlack(picURL)
	}

	return picURL, nil
}

func pgnImportRetrieveHTML(pgn string) ([]byte, error) {
	data := url.Values{
		"pgn": {pgn},
	}

	resp, err := http.PostForm(apiGameImport, data)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	respHTML, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return respHTML, nil
}

func blackMove(pgn string) (bool, error) {
	if pgn == "" {
		return false, errors.New("wrong pgn format")
	}

	split := strings.Split(pgn, " ")
	return len(split)%2 != 0, nil
}

func getPositionURL(respHTML []byte) string {
	doc, err := goquery.NewDocumentFromReader(bytes.NewBuffer(respHTML))
	if err != nil {
		return ""
	}

	url := ""
	doc.Find("a").Each(func(i int, s *goquery.Selection) {
		v, ok := s.Attr("href")
		if !ok {
			return
		}

		if strings.Contains(v, exportFenGif) {
			url = v
		}
	})

	return url
}

func picUrlWithColorBlack(url string) string {
	if strings.Contains(url, colorWhite) {
		return strings.Replace(url, colorWhite, colorBlack, 1)
	}

	return url + colorBlack
}
