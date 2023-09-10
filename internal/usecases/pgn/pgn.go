package pgn

import (
	"bytes"
	"errors"
	"github.com/PuerkitoBio/goquery"
	"io"
	"net/http"
	"net/url"
)

const apiGameImport = "https://lichess.org/api/import"

func GetPictureURL(pgn string) (string, error) {
	respHTML, err := pgnImportRetrieveHTML(pgn)
	if err != nil {
		return "", err
	}

	doc, err := goquery.NewDocumentFromReader(bytes.NewBuffer(respHTML))
	pic := doc.Find(".text.position-gif")

	picURL, exists := pic.Attr("href")
	if !exists {
		return "", errors.New("picture not found")
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
