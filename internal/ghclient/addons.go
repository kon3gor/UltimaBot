package ghclient

import (
	"io/ioutil"
)

func (c *GithubClient) GetFile(req ContentRequest) (string, error) {
	content, err := c.GetContent(req)
	if err != nil {
		return "", err
	}

	res, err := c.httpClient.Get(content[0].DownloadUrl)
	if err != nil {
		return "", err
	}

	defer res.Body.Close()

	bodyBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	return string(bodyBytes), nil
}
