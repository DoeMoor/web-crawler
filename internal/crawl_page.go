package internal

import (
	"fmt"
	"net/url"
)

func CrawlPage(rawBaseURL, rawCurrentURL string, pages map[string]int) {
	if !checkDomain(rawBaseURL, rawCurrentURL) {
		return
	}

	normalizedCurrentURL, err := normalizeURL(rawCurrentURL)
	if err != nil {
		fmt.Println(err)
		return
	}

	if pages[normalizedCurrentURL] > 0 {
		pages[normalizedCurrentURL]++
		return
	}
	
	fmt.Println("Crawling: ", normalizedCurrentURL)
	
	pageHTML, err := getHTML(rawCurrentURL)
	if err != nil {
		fmt.Println(err)
		return
	}
	
	urls, err := getURLFromHTML(pageHTML, rawCurrentURL)
	if err != nil {
		fmt.Println(err)
		return
	}

	pages[normalizedCurrentURL]++
	
	for _, url := range urls {
		CrawlPage(rawBaseURL, url, pages)
	}

	fmt.Println("Crawled: ", normalizedCurrentURL)
}

// checkDomain checks if the base URL and the current URL have the same domain.
//	If they don't, it returns false.
func checkDomain(baseURL, currentURL string) bool {
	baseDomain, err := url.Parse(baseURL)
	if err != nil {
		fmt.Println("checkDomain \"baseURL\" err: ",err)
		return false 
	}

	currentDomain, err := url.Parse(currentURL)
	if err != nil {
		fmt.Println("checkDomain \"currentUrl\" err: ",err)
		return false 
	}

	return baseDomain.Host == currentDomain.Host
}