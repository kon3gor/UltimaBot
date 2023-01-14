package ghclient

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httputil"
	"os"
)

const BASE_URL = "https://api.github.com"

func get(client *http.Client, url string) (*http.Response, error) {
	furl := makeUrl(url)
	request, err := http.NewRequest(http.MethodGet, furl, nil)
	if err != nil {
		return nil, err
	}
	fillRequestWitheaders(request, make(map[string]string))
	return ghrequest(client, request)
}

func post(client *http.Client, url string, body io.Reader) (*http.Response, error) {
	req, err := jsonBodyRequest(url, body, http.MethodPost)
	if err != nil {
		return nil, err
	}
	return ghrequest(client, req)
}

func patch(client *http.Client, url string, body io.Reader) (*http.Response, error) {
	req, err := jsonBodyRequest(url, body, http.MethodPatch)
	if err != nil {
		return nil, err
	}
	return ghrequest(client, req)
}

func jsonBodyRequest(url string, body io.Reader, method string) (*http.Request, error) {
	furl := makeUrl(url)
	headers := map[string]string{
		"Content-Type": "application/json",
	}
	request, err := http.NewRequest(method, furl, body)
	if err != nil {
		return nil, err
	}
	fillRequestWitheaders(request, headers)
	return request, nil
}

func ghrequest(client *http.Client, request *http.Request) (*http.Response, error) {
	d, _ := httputil.DumpRequestOut(request, true)
	fmt.Printf("%q", d)

	res, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	return res, err
}

func makeUrl(url string) string {
	return fmt.Sprintf("%s/%s", BASE_URL, url)
}

func fillRequestWitheaders(req *http.Request, headers map[string]string) {
	baseHeaders := getBaseGithubHeaders()
	for key, value := range baseHeaders {
		req.Header.Set(key, value)
	}

	for key, value := range headers {
		req.Header.Set(key, value)
	}
}

func getBaseGithubHeaders() map[string]string {
	githubToken := os.Getenv("GITHUB_TOKEN")
	return map[string]string{
		"Accept":        "application/vnd.github+json",
		"Authorization": fmt.Sprintf("Bearer %s", githubToken),
	}
}
