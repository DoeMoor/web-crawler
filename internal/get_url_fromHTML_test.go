package internal

import (
	"reflect"
	"testing"
)

func Test_getURLFromHTML(t *testing.T) {
	test := []struct {
		name       string
		rawBaseURL string
		htmlBody   string
		expected   []string
	}{
		{
			name:       "no links",
			rawBaseURL: "http://example.com",
			htmlBody: `
<html>
  <body>
	</body>
</html>`,
			expected: []string{},
		},

		{
			name:       "one relative link",
			rawBaseURL: "http://example.com/st",
			htmlBody: `
<html>
	<body>
		<a href="/st/1">link</a>
	</body>
</html>`,
			expected: []string{"http://example.com/st/st/1"},
		},

		{
			name:       "relative + absolute link",
			rawBaseURL: "http://example.com/st/",
			htmlBody: `
<html>
	<body>
		<a href="/link/st/1">link</a>
		<a href="http://example.com/">link</a>
	</body>
</html>`,
			expected: []string{"http://example.com/st/link/st/1", "http://example.com"},
		},

		{
			name:       "relative link 2 + absolute link 1",
			rawBaseURL: "http://example.com/st",
			htmlBody: `
<html>
	<body>
		<a href="/link/st/1">link</a>
		<a href="/link/st/2">link</a>
	</body>
	<a href="http://example.com/st/i">link</a>
	<a href="hide?id=41546769&amp;goto=news">hide</a>
</html>`,
			expected: []string{"http://example.com/st/link/st/1", "http://example.com/st/link/st/2", "http://example.com/st/i","http://example.com/st/hide?id=41546769&goto=news"},
		},
	}

	for _, tc := range test {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := getURLFromHTML(tc.htmlBody, tc.rawBaseURL)
			if err != nil {
				t.Errorf("unexpected error: %v", err)
				return
			}

			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("expected %v, got %v", tc.expected, actual)
				return
			}

			// t.Logf("test pass %v: \n -- expected \"%v\" \n -- actual   \"%v\"", tc.name,tc.expected, actual)
		})

	}
}
