package httphelper

import (
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

//RequestTheURL makes a GET request against the
//given raw URL. It only supports those URLs
//which returns "text/html" in their response.
func RequestTheURL(rawURL string) (*http.Response, error) {
	resp, err := http.Get(rawURL)
	contentType := resp.Header.Get("Content-Type")
	if !strings.HasPrefix(contentType, "text/html") {
		errMsg := "Incompatible Mime Type. Expecting text/html got "
		return nil, errors.New(errMsg + resp.Header.Get("Content-Type"))
	}
	return resp, err
}

//HTTPGetRequester implements the Requester interface
type HTTPGetRequester struct {
	URL string
}

//Request is the HTTP Get Concrete implementation of the
//Requester interface.
func (req HTTPGetRequester) Request() ([]byte, error) {
	log.Println("HTTP Request called for: ", req.URL)
	if req.URL == "" {
		return nil, errors.New("Illegal State: No URL to crawl")
	}
	resp, err := http.Get(req.URL)
	if err != nil {
		return nil, err
	}
	contentType := resp.Header.Get("Content-Type")
	if !strings.HasPrefix(contentType, "text/html") {
		errMsg := "Incompatible Mime Type. Expecting text/html got "
		return nil, errors.New(errMsg + resp.Header.Get("Content-Type"))
	}
	dataSl, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return dataSl, nil
}
