package main

import (
	"errors"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/nawazish-github/web-crawler/hrefhandler"

	"github.com/nawazish-github/web-crawler/parsers"
)

const depth = 1

var (
	isFirstItr = true
	urlReg     = make(map[string][]string)
)

func main() {
	rawURL := os.Args[1]
	for i := 0; i < depth; i++ {
		//allow first iteration to execute without any checks.
		if !isFirstItr && len(urlReg[rawURL]) == 0 {
			break
		}
		pURL, urlParseErr := parsers.ParseURL(rawURL)
		if urlParseErr != nil {
			log.Fatal("URL Parse Error: ", urlParseErr)
			continue
		}
		isFirstItr = false
		addURLToURLRegistry(rawURL)
		resp, err := requestTheURL(rawURL)
		if err != nil {
			log.Fatal("Request Failure: ", err)
			continue
		}
		defer resp.Body.Close()
		rootElem, htmlParseErr := parsers.ParseHTMLDoc(resp.Body)
		if htmlParseErr != nil {
			log.Fatal("Request Failure: ", err)
			continue
		}
		hrefhandler.Handle(rootElem, pURL, rawURL)
	}
	log.Println(hrefhandler.GetURLReg())
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
