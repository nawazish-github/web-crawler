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
	}
	crawler.Depth = depth

	log.Println(rawURL)
	log.Println(crawler.Depth)
	rawURLs := []string{rawURL}
	gens := crawler.MultiCrawler(rawURLs)
	for _, gen := range gens {
		log.Println(gen)
	}
}
