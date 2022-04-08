package model

import (
	browser "github.com/EDDYCJY/fake-useragent"
	"net/http"
)

var UA []string
var length int

func StatusCode(url string) int {
	randomUA := browser.Random()
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("User-Agent", randomUA)

	resp, err := client.Get(url)
	if err != nil {
		return 404
	}
	Status := resp.StatusCode
	return Status
}
