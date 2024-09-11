package main

import (
	"fmt"
	"os"

	"192.168.1.21/doe/web-crawler/internal"
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

		pages := make(map[string]int)
		internal.CrawlPage(argsWithoutProg[0], argsWithoutProg[0], pages)
		fmt.Println("Crawling: ", argsWithoutProg[0])
		
		for page , count := range pages {
			fmt.Printf("Page: %v, Count: %v\n", page, count)
		}
	}
}
