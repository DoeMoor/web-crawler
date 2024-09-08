package internal

import (
	"errors"
	// "fmt"
	"net/url"
	"strings"
)

func NormalizeURL(inURL string) (string, error) {
	parsedURL, err := url.Parse(inURL)

	if err != nil {
		return "", errors.New("invalid URL")
	}

	path := parsedURL.Host + parsedURL.Path
	path = strings.TrimSuffix(path, "/")
	path = strings.ToLower(path)
	return path, nil
}
