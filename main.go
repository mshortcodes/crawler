package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Println("no website provided")
		os.Exit(1)
	}

	if len(args) > 2 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	}

	rawBaseURL := args[1]
	if len(args) == 2 {
		fmt.Printf("starting crawl of: %s\n", rawBaseURL)
	}

	pages := map[string]int{}
	crawlPage(rawBaseURL, rawBaseURL, pages)

	for page, count := range pages {
		fmt.Printf("%s: %d\n", page, count)
	}
}
