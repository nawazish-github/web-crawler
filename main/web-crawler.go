package main

import (
	"errors"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
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
	resp, err := requestTheURL(rawURL)
	if err != nil {
		log.Fatal("Request Failure: ", err)
		return
	}
	defer resp.Body.Close()
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

func requestTheURL(rawURL string) (*http.Response, error) {
	resp, err := http.Get(rawURL)
	contentType := resp.Header.Get("Content-Type")
	if !strings.HasPrefix(contentType, "text/html") {
		errMsg := "Incompatible Mime Type. Expecting text/html got "
		return nil, errors.New(errMsg + resp.Header.Get("Content-Type"))
	}
	return resp, err
}
