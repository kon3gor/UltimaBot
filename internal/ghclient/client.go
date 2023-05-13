package ghclient

import (
	"fmt"
	"net/http"
)

type headers map[string]string

type GithubClient struct {
	baseUrl     string
	baseHeaders headers
	httpClient  *http.Client
}

type feature func(client *GithubClient)

func NewClient(token string, features ...feature) *GithubClient {
	client := defaultClient(token)
	for _, f := range features {
		f(client)
	}
	return client
}

const (
	baseUrl = "https://api.github.com"
	accept  = "application/vnd.github+json"
)

func defaultClient(token string) *GithubClient {
	ftoken := fmt.Sprintf("Bearer %s", token)
	return &GithubClient{
		baseUrl:    baseUrl,
		httpClient: http.DefaultClient,
		baseHeaders: map[string]string{
			"Accept":        accept,
			"Authorization": ftoken,
		},
	}
}

func WithBaseUrl(url string) feature {
	return func(c *GithubClient) {
		c.baseUrl = url
	}
}

func WithBaseHeaders(headers headers) feature {
	return func(c *GithubClient) {
		for k, v := range headers {
			c.baseHeaders[k] = v
		}
	}
}

func WithHttpClient(client *http.Client) feature {
	return func(c *GithubClient) {
		c.httpClient = client
	}
}
