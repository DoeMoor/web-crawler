package internal

import (
	"fmt"
	"net/url"
	"strings"

	"golang.org/x/net/html"
)


func GetURLFromHTML(htmlBody, rawBaseURL string)([]string, error) {

func getURLFromHTML(htmlBody, rawBaseURL string)([]string, error) {

	nodDocument, err := html.Parse(strings.NewReader(htmlBody))
	if err != nil {
		return nil, err
	}

	var URLs []string = []string{}

	var f func(*html.Node, *[]string)
	f = func(nodElement *html.Node, URLs *[]string) {
		if nodElement.Type == html.ElementNode && nodElement.Data == "a" {
			for _, attr := range nodElement.Attr {
				if attr.Key == "href"{
					absoluteURL, err := getAbsoluteURL(rawBaseURL, attr.Val)
					if err != nil {
						fmt.Println("Error: ", err)
						continue
					}
					*URLs = append(*URLs, absoluteURL)
				}
			}
	  }
		for c := nodElement.FirstChild; c != nil; c = c.NextSibling {
			f(c, URLs)
		}
	}
	f(nodDocument, &URLs)

	return URLs, nil
}


func getAbsoluteURL(rawBaseURL, href string) (string, error) {
	hrefURL, err := url.Parse(href)
	if err != nil {
		return "", err
	}

	if hrefURL.IsAbs() {
		return href, nil
	}

	return rawBaseURL + hrefURL.String(), nil
	
}

