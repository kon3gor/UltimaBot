package ghclient

import (
	"io"
	"log"
	"net/http"
	"net/http/httputil"
)

var contentTypeJson = map[string]string{
	"Content-Type": "application/json",
}

func (c *GithubClient) get(url string) (*http.Response, error) {
	furl := c.formatUrl(url)
	req, err := http.NewRequest(http.MethodGet, furl, nil)
	c.fillHeaders(req, nil)
	if err != nil {
		return nil, err
	}
	return c.do(req)
}

func (c *GithubClient) post(url string, body io.Reader) (*http.Response, error) {
	furl := c.formatUrl(url)
	req, err := http.NewRequest(http.MethodPost, furl, body)
	c.fillHeaders(req, contentTypeJson)
	if err != nil {
		return nil, err
	}
	return c.do(req)
}

func (c *GithubClient) patch(url string, body io.Reader) (*http.Response, error) {
	furl := c.formatUrl(url)
	req, err := http.NewRequest(http.MethodPatch, furl, body)
	c.fillHeaders(req, contentTypeJson)
	if err != nil {
		return nil, err
	}
	return c.do(req)
}

func (c *GithubClient) do(req *http.Request) (*http.Response, error) {
	d, _ := httputil.DumpRequestOut(req, true)
	log.Printf("%q", d)
	return c.httpClient.Do(req)
}
