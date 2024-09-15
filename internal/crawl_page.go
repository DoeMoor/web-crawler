package internal

import (
	"fmt"
	"net/url"
	"sync"
)

type config struct {
	Pages              map[string]int
	baseURL            url.URL
	Mu                 sync.Mutex
	concurrencyControl chan struct{}
	Wg                 sync.WaitGroup
	maxPages           int
}

// NewConfig creates a new config for crawler.
//
//	config{
//		Pages:              make(map[string]int),
//		BaseURL:            *baseURL,
//		Mu:                 sync.Mutex{},
//		ConcurrencyControl: make(chan struct{}, maxConcurrency),
//		Wg:                 sync.WaitGroup{},
//		maxPages:           maxPages,
//	}
func NewCrawlerConfig(rawBaseURL string, maxConcurrency, maxPages int) (*config, error) {
	baseURL, err := url.Parse(rawBaseURL)
	if err != nil {
		return nil, err
	}

	return &config{
		Pages:              make(map[string]int),
		baseURL:            *baseURL,
		Mu:                 sync.Mutex{},
		concurrencyControl: make(chan struct{}, maxConcurrency),
		Wg:                 sync.WaitGroup{},
		maxPages:           maxPages,
	}, nil
}

// CrawlPage crawls a page and extracts all URLs from it.
func (cnf *config) CrawlPage(rawCurrentURL string) {

	cnf.acquireConcurrencySlot()

	defer func() {
		cnf.Wg.Done()
		<-cnf.concurrencyControl
	}()

	if cnf.PagesLen() >= cnf.maxPages {
		return
	}

	//skip if the base URL and the current URL have different domains
	if !checkDomain(cnf.baseURL.String(), rawCurrentURL) {
		return
	}

	normalizedCurrentURL, err := normalizeURL(rawCurrentURL)
	if err != nil {
		fmt.Println(err)
		return
	}

	ok := cnf.addVisitedPage(normalizedCurrentURL)
	if !ok {
		return
	}

	//check if the limit of pages is reached

	pageHTML, err := getHTML(rawCurrentURL)
	if err != nil {
		fmt.Println(err)
		return
	}

	urls, err := getURLFromHTML(pageHTML, cnf.baseURL.String())
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, url := range urls {
		if cnf.PagesLen() >= cnf.maxPages {
			break
		}
		cnf.Wg.Add(1)
		go cnf.CrawlPage(url)
	}

	fmt.Println("Crawled: ", normalizedCurrentURL)
}

func (cnf *config) acquireConcurrencySlot() {
	cnf.concurrencyControl <- struct{}{}
}

// isPagesLimitReached checks if the number of pages crawled is equal to the maximum number of pages.
// if maxPages is reached, it returns 	"true".
// otherwise, it returns 	"false".
func (cnf *config) PagesLen() int {
	defer cnf.Mu.Unlock()
	cnf.Mu.Lock()
	return len(cnf.Pages)
}

// checkDomain checks if the base URL and the current URL have the same domain.
//
// If they don't, it returns "false".
func checkDomain(baseURL, currentURL string) bool {
	baseDomain, err := url.Parse(baseURL)
	if err != nil {
		fmt.Println("checkDomain \"baseURL\" err: ", err)
		return false
	}

	currentDomain, err := url.Parse(currentURL)
	if err != nil {
		fmt.Println("checkDomain \"currentUrl\" err: ", err)
		return false
	}

	return baseDomain.Host == currentDomain.Host
}

// If page is added to the map, it returns "true",
//
// otherwise it increments the page count and returns "false".
func (cnf *config) addVisitedPage(normalizedCurrentURL string) bool {

	cnf.Mu.Lock()
	defer cnf.Mu.Unlock()

	_, pageIsVisited := cnf.Pages[normalizedCurrentURL]
	if pageIsVisited {
		cnf.Pages[normalizedCurrentURL]++ //increment the page count
		return false
	}
	cnf.Pages[normalizedCurrentURL] = 1
	return true //page is added to the map
}
