package ghclient

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type FileContent struct {
	DownloadUrl string `json:"download_url"`
}

type ContentRequest struct {
	path string
}

func GetContent(client *http.Client, req ContentRequest) ([]FileContent, error) {
	res, err := get(client, req.path)
	defer res.Body.Close()

	if req.isFile() {
		var filecontent FileContent
		if err = json.NewDecoder(res.Body).Decode(&filecontent); err != nil {
			return nil, err
		}
		return []FileContent{filecontent}, nil
	} else {
		var filescontent []FileContent
		if err = json.NewDecoder(res.Body).Decode(&filescontent); err != nil {
			return nil, err
		}
		return filescontent, nil
	}
}

func NewMyContentRequest(repo string, path string) ContentRequest {
	return NewContentRequest("kon3gor", repo, path)
}

func NewContentRequest(user string, repo string, path string) ContentRequest {
	fullPath := fmt.Sprintf("repos/%s/%s/contents/%s", user, repo, path)
	return ContentRequest{fullPath}
}

// If last segment of the provided path has a period, then it is probably a file.
// Also we should check last period index, because hidden files and dirs start with a period.
func (self ContentRequest) isFile() bool {
	parts := strings.Split(self.path, "/")
	filename := parts[len(parts)-1]
	return strings.Contains(filename, ".") && strings.LastIndex(filename, ".") != 0
}
