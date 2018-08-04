package main

import (
	"errors"
	"log"
	"net/http"
	"net/url"
	"strings"

	"golang.org/x/net/html"
)

var urlReg = make(map[string][]string)

func main() {
	//rawURL := os.Args[1]
	rawURL := "https://twitter.com"
	pURL, err := ParseURL(rawURL)
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
	node, parseErr := html.Parse(resp.Body)
	if parseErr != nil {
		log.Fatal("html parse failure")
		return
	}
	var hrefIterator func(*html.Node)
	hrefIterator = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, atr := range n.Attr {
				if atr.Key == "href" {
					link := atr.Val
					if isRelativeURL(link) {
						link = constructAbsoluteURL(pURL, link)
						updateURLRegistryWithLatestLink(rawURL, link)
						continue
					}
					verifyAndUpdateURLRegistryWithLatestLink(link, rawURL, pURL.Host)
				}
			}
		}
		for elem := n.FirstChild; elem != nil; elem = elem.NextSibling {
			hrefIterator(elem)
		}
	}
	hrefIterator(node)
	log.Println(urlReg)
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

func isRelativeURL(link string) bool {
	if strings.HasPrefix(link, "/") {
		return true
	}
	return false
}

func constructAbsoluteURL(pURL *url.URL, link string) string {
	return pURL.Scheme + "://" + pURL.Host + link
}

func updateURLRegistryWithLatestLink(rawURL, link string) {
	list := urlReg[rawURL]
	list = append(list, link)
	urlReg[rawURL] = list
}

func verifyAndUpdateURLRegistryWithLatestLink(link, rawURL, host string) {
	pURL, err := ParseURL(link)
	if err != nil {
		log.Fatal("Could not parse the link: ", err)
		return
	}
	if !strings.Contains(pURL.Host, host) {
		return
	}
	updateURLRegistryWithLatestLink(rawURL, link)
}
