package crawler

import (
	"log"
	"net/url"
	"strings"

	"github.com/nawazish-github/web-crawler/httphelper"
	"github.com/nawazish-github/web-crawler/model"

	"github.com/nawazish-github/web-crawler/parsers"

	"golang.org/x/net/html"
)

//var urlReg = make(map[string][]string)

//HrefHandler handles all the href tags
//embedded in the given document and
//updates the same to the urlRegistry
var hrefHandler func(*html.Node, *url.URL, string, *model.Generation)

//Crawl iterates the HTML document beginning
//at its root and updates the URL registry with
//each encountered link.
func Crawl(rootElem *html.Node, pURL *url.URL, rawURL string, gen *model.Generation) {
	hrefHandler = func(n *html.Node, pURL *url.URL, rawURL string, gen *model.Generation) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, atr := range n.Attr {
				if atr.Key == "href" {
					link := atr.Val
					if isRelativeURL(link) {
						link = constructAbsoluteURL(pURL, link)
						updateURLRegistryWithLatestLink(rawURL, link, gen)
						continue
					}
					verifyAndUpdateURLRegistryWithLatestLink(link, rawURL, pURL.Host, gen)
				}
			}
		}
		for elem := n.FirstChild; elem != nil; elem = elem.NextSibling {
			hrefHandler(elem, pURL, rawURL, gen)
		}
	}
	hrefHandler(rootElem, pURL, rawURL, gen)
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

func updateURLRegistryWithLatestLink(rawURL, link string, gen *model.Generation) {
	list := gen.GenMap[rawURL]
	list = append(list, link)
	gen.GenMap[rawURL] = list
}

//verifyAndUpdateURLRegistryWithLatestLink checks if the link under
//consideration is from the same base domain; it discards any
//external links.
func verifyAndUpdateURLRegistryWithLatestLink(link, rawURL, host string, gen *model.Generation) {
	pURL, err := parsers.ParseURL(link)
	if err != nil {
		log.Fatal("Could not parse the link: ", err)
		return
	}
	if !strings.Contains(pURL.Host, host) {
		return
	}
	updateURLRegistryWithLatestLink(rawURL, link, gen)
}

var siteMap []model.Generation
var depth = 1
var counter = 0

//MultiCrawler crawls a list of URLs and returns a list of
//generations.
func MultiCrawler(rawURLs []string) []model.Generation {
	counter++
	if counter > depth {
		return nil
	}
	gen := model.NewGeneration()
	for _, rawURL := range rawURLs {
		pURL, urlParseErr := parsers.ParseURL(rawURL)
		if urlParseErr != nil {
			log.Fatal("URL Parse Error: ", urlParseErr)
			continue
		}
		//gen := model.NewGeneration()
		gen.GenMap[rawURL] = []string{}
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
		Crawl(rootElem, pURL, rawURL, gen)
		siteMap = append(siteMap, *gen)
	}
	for _, rawURLs := range gen.GenMap {
		if len(rawURLs) > 0 {
			MultiCrawler(rawURLs)
		}
	}
	return siteMap
}
