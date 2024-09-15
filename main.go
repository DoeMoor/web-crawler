package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"192.168.1.21/doe/web-crawler/internal"
)

func main() {

	// Gents from the command line
	args := os.Args[1:]

	switch {
	case len(args) == 0:
		fmt.Println("usage: <url> [concurrency control] [max pages]",
			"\"example: http://example.com 5 10\"")
		os.Exit(0)
	case len(args) > 5:
		fmt.Println("too many arguments")
		os.Exit(1)
	case len(args) >= 1:
		startCrawl(args)
	}

}

func startCrawl(args []string) {
	var maxConcurrency int = 5
	var maxPages int = 10
	var rawBaseURL string = strings.TrimSpace(args[0])

	// Check if the user has provided the concurrency control and max pages
	switch len(args) {
	case 1:
		// Do nothing

	case 2: // Check if the user has provided the concurrency control
		arg1, err := strconv.Atoi(args[1])
		if err != nil {
			fmt.Printf(`
->error parsing arguments:
		using values
			concurrency control: %v  -- user provided: %v
			max pages:           %v`, maxConcurrency, args[1], maxPages)
			fmt.Println()
			fmt.Println()
			return
		}
		maxConcurrency = arg1
	default: // Check if the user has provided the concurrency control and max pages
		arg, err := strconv.Atoi(args[1]) // Check if the user has provided correct the concurrency control
		if err == nil {
			maxConcurrency = arg
		}

		arg, err = strconv.Atoi(args[2]) // Check if the user has provided correct the max pages
		if err == nil {
			maxPages = arg
		}

		if err != nil {
			fmt.Printf(`
->error parsing arguments:
		using values
			concurrency control: %v  -- user provided: %v
			max pages:           %v  -- user provided: %v`, maxConcurrency, args[1], maxPages, args[2])
			fmt.Println()
			fmt.Println()
		}
	}

	// fmt.Printf("starting crawl of: \"%v\" %v %v\n", rawBaseURL, maxConcurrency, maxPages)

	crawler, err := internal.NewCrawlerConfig(rawBaseURL, maxConcurrency, maxPages)
	if err != nil {
		fmt.Println("error creating config")
		os.Exit(1)
	}

	crawler.Wg.Add(1)
	go crawler.CrawlPage(rawBaseURL)
	crawler.Wg.Wait()

	// fmt.Println("=================== all pages crawled ===================")
	// fmt.Printf( "=================== Pages crawled: %v ===================\n", crawler.PagesLen())
	crawler.PrintReport()
	// fmt.Println("=========================================================")
}
