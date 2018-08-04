package main

import (
	"log"
	"net/url"
	"os"
)

func main() {
	rawURL := os.Args[1]
	_, err := ParseURL(rawURL)
	if err != nil {
		log.Fatal("Parse failure")
		return
	}
}

//ParseURL parses a raw URL
func ParseURL(rawURL string) (*url.URL, error) {
	pURL, err := url.Parse(rawURL)
	return pURL, err
}
