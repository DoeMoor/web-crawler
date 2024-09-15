package main

import (
	"192.168.1.21/doe/web-crawler/internal"
	"fmt"
	"os"
)

func main() {
	argsWithoutProg := os.Args[1:]

	if len(argsWithoutProg) < 1 {
		fmt.Println("no website provided")
		os.Exit(1)
	}
	if len(argsWithoutProg) > 1 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	}
	if len(argsWithoutProg) == 1 {
		fmt.Printf("starting crawl of: \"%v\"\n", argsWithoutProg[0])

		rawBaseURL := argsWithoutProg[0]

		crawler, err := internal.NewCrawlerConfig(rawBaseURL, 50, 3000)
		if err != nil {
			fmt.Println("error creating config")
			os.Exit(1)
		}

		crawler.Wg.Add(1)
		go crawler.CrawlPage(rawBaseURL)
		crawler.Wg.Wait()
		
		fmt.Println("=================== all pages crawled ===================")

		crawler.Mu.Lock()
		for page, count := range crawler.Pages {
			fmt.Printf("Page: %v, Count: %v\n", page, count)
		}
		crawler.Mu.Unlock()
	}
}
