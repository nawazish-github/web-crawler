package main

import (
	"os"

	"github.com/nawazish-github/web-crawler/crawler"
)

func main() {
	rawURL := os.Args[1]
	rawURLs := []string{rawURL}
	crawler.MultiCrawler(rawURLs)
}
