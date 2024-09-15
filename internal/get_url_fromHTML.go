package internal

import (
	"fmt"
	"net/url"
	"strings"

	"golang.org/x/net/html"
)

// getURLFromHTML extracts all URLs from an HTML document.
//
//	It returns a slice of URLs.
//	e.g. getURLFromHTML("<a href='http://example.com/st/ss/1'>", "http://example.com/st/ss/1")
func getURLFromHTML(htmlBody, rawBaseURL string) ([]string, error) {

	nodDocument, err := html.Parse(strings.NewReader(htmlBody))
	if err != nil {
		return nil, err
	}

	var URLs []string = []string{}

	var getNode func(*html.Node)
	getNode = func(nodElement *html.Node) {
		if nodElement.Type == html.ElementNode && nodElement.Data == "a" {
			for _, attr := range nodElement.Attr {
				if attr.Key == "href" {
					absoluteURL, err := getAbsoluteURL(rawBaseURL, attr.Val)
					if err != nil {
						fmt.Println("Error: ", err)
						continue
					}
					if absoluteURL != rawBaseURL {
						URLs = append(URLs, absoluteURL)
					}
				}
			}
		}
		for child := nodElement.FirstChild; child != nil; child = child.NextSibling {
			getNode(child)
		}
	}

	getNode(nodDocument)

	return URLs, nil
}

// getAbsoluteURL returns the absolute URL from a base URL and a relative URL.
//
//	e.g. getAbsoluteURL("http://example.com/st", "/ss/1") => "http://example.com/st/ss/1"
func getAbsoluteURL(rawBaseURL, href string) (string, error) {
	rawBaseURL = strings.TrimSuffix(rawBaseURL, "/")
	href = strings.TrimSuffix(href, "/")

	hrefURL, err := url.Parse(href)
	if err != nil {
		return "", err
	}

	if hrefURL.IsAbs() {

		return href, nil
	}

	absoluteURL, err := url.JoinPath(rawBaseURL, href)
	if err != nil {
		return "", err
	}

	absoluteURL, err = url.PathUnescape(absoluteURL)
	if err != nil {
		return "", err
	}
	

	return absoluteURL, nil

}
