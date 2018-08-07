package parsers

import (
	"bytes"
	"net/url"

	"golang.org/x/net/html"
)

//ParseURL parses a raw URL
func ParseURL(rawURL string) (*url.URL, error) {
	pURL, err := url.Parse(rawURL)
	return pURL, err
}

//ParseHTMLDoc parses given argument into a valid
//HTML doc returning its root element, Node or error
func ParseHTMLDoc(data []byte) (*html.Node, error) {
	doc := bytes.NewReader(data)
	return html.Parse(doc)
}
