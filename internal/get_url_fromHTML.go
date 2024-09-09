package internal

import (
	"fmt"
	"net/url"
	"strings"

	"golang.org/x/net/html"
)

func getURLFromHTML(htmlBody, rawBaseURL string)([]string, error) {

	nodDocument, err := html.Parse(strings.NewReader(htmlBody))
	if err != nil {
		return nil, err
	}

	var URLs []string = []string{}

	var getNode func(*html.Node)
	getNode = func(nodElement *html.Node) {
		if nodElement.Type == html.ElementNode && nodElement.Data == "a" {
			for _, attr := range nodElement.Attr {
				if attr.Key == "href"{
					absoluteURL, err := getAbsoluteURL(rawBaseURL, attr.Val)
					if err != nil {
						fmt.Println("Error: ", err)
						continue
					}
					URLs = append(URLs, absoluteURL)
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

