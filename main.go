package main

import (
	"fmt"
	// "net/http"
	"os"
)

func main() {
	// argsWithProg := os.Args
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
	}
	// fmt.Println(argsWithProg)
	// fmt.Println(argsWithoutProg)

	// argsURL := argsWithoutProg[0]

	// client := http.Client{}
	// request, err := http.NewRequest("GET", argsURL, nil)
	// if err != nil {
	// 	fmt.Println("new request error: ", err,"\n", argsURL)
	// 	return
	// }

	// respond, err := client.Do(request)
	// if err != nil {
	// 	fmt.Println("request error: ", err,"\n", request)
	// 	return
	// }

}
