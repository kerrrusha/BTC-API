package utils

import (
	"io"
	"net/http"
	"time"
)

func RequestJson(url string) []byte {
	client := http.Client{Timeout: time.Second * 2}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	CheckForError(err)

	res, err := client.Do(req)
	CheckForError(err)

	if res.Body != nil {
		defer func(Body io.ReadCloser) {
			err := Body.Close()
			CheckForError(err)
		}(res.Body)
	}

	body, readErr := io.ReadAll(res.Body)
	CheckForError(readErr)

	return body
}
