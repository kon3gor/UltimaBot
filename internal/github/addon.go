package github

import (
	"dev/kon3gor/ultima/pkg/ghclient"
	"io/ioutil"
	"net/http"
	"os"
)

var client = ghclient.NewClient(os.Getenv("GITHUB_TOKEN"))

const (
	usernmae = "kon3gor"
	repo     = "PersonalObsidian"
)

func GetObsidianFile(path string) (string, error) {
	req := ghclient.NewContentRequest(usernmae, repo, path)
	content, err := client.GetContent(req)
	if err != nil {
		return "", err
	}

	res, err := http.Get(content[0].DownloadUrl)
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

func GetObsidianFolder(path string) ([]ghclient.FileContent, error) {
	req := ghclient.NewContentRequest(usernmae, repo, path)
	content, err := client.GetContent(req)
	if err != nil {
		return nil, err
	}

	return content, nil
}

func SaveObisdianFile(path string, content string) error {
	req := ghclient.NewPushRequest(usernmae, repo, "main", path, content)
	return client.PushContent(req)
}
