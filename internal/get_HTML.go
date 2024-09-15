package internal

import (
	// "errors"
	"fmt"
	"io"
	"net/http"
)

// getHTML makes a GET request to the URL and returns the body of the response as string.
func getHTML(rawURL string) (string, error) {
	client := http.Client{}

	request, err := http.NewRequest("GET", rawURL, nil)
	if err != nil {
		return "", err
	}
	request.Header.Set("Content-Type", "text/html")

	response, err := client.Do(request)
	if err != nil {
		return "", err
	}

	defer response.Body.Close()

	if response.StatusCode >= 400 {
		return "", fmt.Errorf("status code: %v  url: \"%v\"", response.StatusCode, rawURL)
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}
