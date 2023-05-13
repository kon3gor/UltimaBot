package ghclient

import (
	"fmt"
	"net/http"
	"strings"
)

func (c *GithubClient) fillHeaders(req *http.Request, headers headers) {
	for key, value := range c.baseHeaders {
		req.Header.Set(key, value)
	}
	if headers == nil {
		return
	}
	for key, value := range headers {
		req.Header.Set(key, value)
	}
}

func (c *GithubClient) formatUrl(path string) string {
	path = strings.Trim(path, "/")
	return fmt.Sprintf("%s/%s", c.baseUrl, path)
}
