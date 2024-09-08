package internal

import (
	"errors"
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
	return path, nil
}