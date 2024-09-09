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

		bodyString, err := internal.GetHTML(argsWithoutProg[0])
		if err != nil {
			fmt.Println("Error: ", err)
			os.Exit(1)
		}
		fmt.Println(bodyString)
	}
}
