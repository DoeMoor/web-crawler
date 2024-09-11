package internal

import (
	"errors"
	"net/url"
	"strings"
)

// NormalizeURL removes the scheme, capital letters, and trailing slash from a URL.
//	inputURL: "http://example.com/some/long/Path/",
//	expected: "example.com/some/long/path",
func normalizeURL(inURL string) (string, error) {
	parsedURL, err := url.Parse(inURL)

	if err != nil {
		return "", errors.New("invalid URL")
	}

	path := parsedURL.Host + parsedURL.Path
	path = strings.TrimSuffix(path, "/")
	path = strings.ToLower(path)
	return path, nil
}
