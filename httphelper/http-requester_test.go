package httphelper

import (
	"errors"
	"testing"
)

func TestShouldReturnValidHTMLResponseWhenURLProvided(t *testing.T) {
	httpReq := StubHTTPGetRequester{URL: "https//twitter.com", resp: []byte("<html></html>")}
	resp, err := httpReq.Request()
	if err != nil {
		t.Errorf("HTTP Get request failed!")
		return
	}
	if string(resp) != "<html></html>" {
		t.Errorf("HTTP Get did prematurely returned!")
		return
	}
}

func TestShouldReturnErrorWhenURLNotProvided(t *testing.T) {
	stubbedResp := []byte("<html></html>")
	httpReq := StubHTTPGetRequester{URL: "", resp: stubbedResp}
	resp, err := httpReq.Request()
	if err == nil {
		t.Errorf("HTTP Get did not return error when no url was provided!")
		return
	}
	if len(resp) == len(stubbedResp) {
		t.Errorf("HTTP Get returned full response when no url was provided!")
		return
	}
}

type StubHTTPGetRequester struct {
	URL  string
	resp []byte
}

func (httpStub StubHTTPGetRequester) Request() ([]byte, error) {
	if httpStub.URL != "https//twitter.com" || httpStub.URL == "" {
		return nil, errors.New("Illegal State: No URL to crawl")
	}
	return []byte("<html></html>"), nil
}
