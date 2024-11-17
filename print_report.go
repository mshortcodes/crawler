package main

import (
	"fmt"
	"sort"
)

type pageInfo struct {
	count   int
	baseURL string
}

func printReport(pages map[string]int, baseURL string) {
	fmt.Println("==========")
	fmt.Printf("REPORT for %s\n", baseURL)
	fmt.Println("==========")

	sortedPages := sortPages(pages)

	for _, v := range sortedPages {
		fmt.Printf("Found %d internal links to %s\n", v.count, v.baseURL)
	}
}

func sortPages(pages map[string]int) []pageInfo {
	pagesInfo := make([]pageInfo, len(pages))

	for k, v := range pages {
		pagesInfo = append(pagesInfo, pageInfo{
			count:   v,
			baseURL: k,
		})
	}

	sort.Slice(pagesInfo, func(i, j int) bool {
		return pagesInfo[i].count > pagesInfo[j].count
	})

	return pagesInfo
}
