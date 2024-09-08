package internal

import (
	"fmt"
	"testing"
)

func TestNormalizeURL(t *testing.T) {
	test := []struct {
		name     string
		inputURL string
		expected string
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
	}

	var passed []string

	for i, tc := range test {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := NormalizeURL(tc.inputURL)
			if err != nil {
				t.Errorf("test %v: unexpected error: %v", i, err)
			}
			if actual != tc.expected {
				t.Errorf("test %v: expected %v, got %v", i, tc.expected, actual)
			}
			passed = append(passed, fmt.Sprintf("test %v: %v", i, tc.name))
		})
	}

	fmt.Println("Passed tests:")
	for test := range passed {
		fmt.Println(passed[test])
	}
}