package crawler

import (
	"log"
	"net/url"
	"strings"

	"github.com/nawazish-github/web-crawler/model"

	"github.com/nawazish-github/web-crawler/parsers"

	"golang.org/x/net/html"
)

//var urlReg = make(map[string][]string)

//HrefHandler handles all the href tags
//embedded in the given document and
//updates the same to the urlRegistry
var hrefHandler func(*html.Node, *url.URL, string)

//Crawl iterates the HTML document beginning
//at its root and updates the URL registry with
//each encountered link.
func Crawl(rootElem *html.Node, pURL *url.URL, rawURL string) {
	hrefHandler = func(n *html.Node, pURL *url.URL, rawURL string) {
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
			hrefHandler(elem, pURL, rawURL)
		}
	}
	hrefHandler(rootElem, pURL, rawURL)
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
	list := model.GetURLReg()[rawURL]
	list = append(list, link)
	model.GetURLReg()[rawURL] = list
	// list := urlReg[rawURL]
	// list = append(list, link)
	// urlReg[rawURL] = list
}

//verifyAndUpdateURLRegistryWithLatestLink checks if the link under
//consideration is from the same base domain; it discards any
//external links.
func verifyAndUpdateURLRegistryWithLatestLink(link, rawURL, host string) {
	pURL, err := parsers.ParseURL(link)
	if err != nil {
		log.Fatal("Could not parse the link: ", err)
		return
	}
	if !strings.Contains(pURL.Host, host) {
		return
	}
	updateURLRegistryWithLatestLink(rawURL, link)
}

//GetURLReg returns the urlReg
// func GetURLReg() map[string][]string {
// 	return urlReg
// }
