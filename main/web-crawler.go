package main

import (
	"log"
	"net/url"
	"os"
)

var urlReg = make(map[string][]string)

func main() {
	rawURL := os.Args[1]
	_, err := ParseURL(rawURL)
	if err != nil {
		log.Fatal("Parse failure")
		return
	}
	addURLToURLRegistry(rawURL)
}

//ParseURL parses a raw URL
func ParseURL(rawURL string) (*url.URL, error) {
	pURL, err := url.Parse(rawURL)
	return pURL, err
}

//addURLToURLRegistry maintains a registry of crawled URLs
func addURLToURLRegistry(rawURL string) {
	if _, ok := urlReg[rawURL]; !ok {
		urlReg[rawURL] = []string{}
	}
}
