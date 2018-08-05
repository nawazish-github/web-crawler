package parsers

import (
	"io"
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
func ParseHTMLDoc(doc io.Reader) (*html.Node, error) {
	return html.Parse(doc)
}
