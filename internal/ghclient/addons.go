package ghclient

import (
	"io/ioutil"
	"net/http"
)

func GetFile(client *http.Client, req ContentRequest) (string, error) {
	content, err := GetContent(client, req)
	if err != nil {
		return "", err
	}

	res, err := client.Get(content[0].DownloadUrl)
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
