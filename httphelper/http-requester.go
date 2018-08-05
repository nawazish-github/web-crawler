package httphelper

import (
	"errors"
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
