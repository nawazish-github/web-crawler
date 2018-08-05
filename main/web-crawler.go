package main

import (
	"log"
	"os"

	"github.com/nawazish-github/web-crawler/httphelper"

	"github.com/nawazish-github/web-crawler/model"

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
		model.AddURLToURLRegistry(rawURL)
		resp, err := httphelper.RequestTheURL(rawURL)
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
	log.Println(model.GetURLReg())
}
