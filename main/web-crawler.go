package main

import (
	"log"
	"os"
	"strconv"

	"github.com/nawazish-github/web-crawler/crawler"
)

func main() {
	rawURL := os.Args[1]
	depth, err := strconv.Atoi(os.Args[len(os.Args)-1])
	if err != nil {
		crawler.Depth = 1
	} else if depth == 0 {
		crawler.Depth = 1
	} else {
		crawler.Depth = depth
	}
	log.Println("Base URL to crawl: ", rawURL)
	log.Println("Depth of crawl: ", crawler.Depth)
	rawURLs := []string{rawURL}
	gens := crawler.MultiCrawler(rawURLs)
	for _, gen := range gens {
		log.Println(gen)
	}
}
