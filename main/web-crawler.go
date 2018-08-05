package main

import (
	"os"

	"github.com/nawazish-github/web-crawler/crawler"
)

var (
	isFirstItr = true
	urlReg     = make(map[string][]string)
)

func main() {
	rawURL := os.Args[1]
	rawURLs := []string{rawURL}
	crawler.MultiCrawler(rawURLs)
}
