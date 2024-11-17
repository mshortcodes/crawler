package main

import (
	"fmt"
	"net/url"
	"os"
	"strconv"
	"sync"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Println("no website provided")
		os.Exit(1)
	}

	if len(args) > 4 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	}

	rawBaseURL := args[1]
	maxConcurrency, _ := strconv.Atoi(args[2])
	maxPages, _ := strconv.Atoi(args[3])

	baseURL, err := url.Parse(rawBaseURL)
	if err != nil {
		fmt.Printf("couldn't parse baseURL: %v", err)
		os.Exit(1)
	}

	cfg := config{
		pages:              make(map[string]int),
		baseURL:            baseURL,
		mu:                 &sync.Mutex{},
		concurrencyControl: make(chan struct{}, maxConcurrency),
		wg:                 &sync.WaitGroup{},
		maxPages:           maxPages,
	}

	fmt.Printf("starting crawl of: %s\n", rawBaseURL)

	cfg.wg.Add(1)
	go cfg.crawlPage(rawBaseURL)
	cfg.wg.Wait()

	for page, count := range cfg.pages {
		fmt.Printf("%s: %d\n", page, count)
	}

	printReport(cfg.pages, rawBaseURL)
}
