package fetcher

import (
	"fmt"
	"net/http"
)

type Fetcher interface {
	Get(string) (*http.Response, error)
}

func (httpFetcher *HttpFetcher) Get(query string) (*http.Response, error) {
	url := fmt.Sprintf("%s://%s:%d/%s", httpFetcher.Protocol, httpFetcher.Host, httpFetcher.Port, query)
	return http.Get(url)
}

type HttpFetcher struct {
	Host     string
	Port     int
	Protocol string
}

func ServiceCheck(query string, fetcher Fetcher) {
	resp, err := fetcher.Get(query)
	if err != nil {
		// todo: do something record
		fmt.Printf("%v", err)
	}

	if resp.StatusCode == 404 {
		// todo: do something record
		fmt.Println("status code is 404")
		return
	} else if resp.StatusCode == 400 {
		// todo: do something record
		fmt.Println("status code is 400")
		return
	} else if resp.StatusCode != 200 {
		// todo: do something record
		return
	}

	// todo: do something record for status code 200
	return
}
