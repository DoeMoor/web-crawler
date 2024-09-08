package internal

import (
	// "fmt"
	"testing"
	"strings"
)

func TestNormalizeURL(t *testing.T) {
	test := []struct {
		name     string
		inputURL string
		expected string
		errorContains string
	}{
		{
			name:     "remove scheme",
			inputURL: "http://example.com",
			expected: "example.com",
		},
		{
			name:     "remove slash",
			inputURL: "http://example.com/",
			expected: "example.com",
		},
		{
			name:     "long url",
			inputURL: "http://example.com/some/long/path",
			expected: "example.com/some/long/path",
		},
		{
			name:     "URL with port",
			inputURL: "http://example.com:8080/rstt/rrr/",
			expected: "example.com:8080/rstt/rrr",
		},
		{
			name:     "URL with query",
			inputURL: "http://example.com/stf/rrr/rstp?query=1",
			expected: "example.com/stf/rrr/rstp",
		},
		{
			name:     "wrong URL",
			inputURL: "://example.com:/rst",
			expected: "",
			errorContains: "invalid URL",
		},
		{
			name:     "remove scheme and capitals and trailing slash",
			inputURL: "http://BLOG.boot.dev/path/",
			expected: "blog.boot.dev/path",
		},
		{
			name:          "handle invalid URL",
			inputURL:      `:\\invalidURL`,
			expected:      "",
			errorContains: "invalid URL",
		},
		{
			name: "path only",
			inputURL: "/stf/rrr/rstp?query=1",
			expected: "/stf/rrr/rstp",
		},
	}

	for i, tc := range test {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := NormalizeURL(tc.inputURL)

			if err != nil && strings.Contains(err.Error(), tc.errorContains) {
				t.Logf("test %v: %v", i, tc.name)
				return
			}
			if err != nil {
				t.Errorf("test %v: unexpected error: %v", i, err)
				return
			}
			if actual != tc.expected {
				t.Errorf("test %v: expected %v, got %v", i, tc.expected, actual)
				return
			}

			t.Logf("test %v: %v", i, tc.name)
		})
	}
}