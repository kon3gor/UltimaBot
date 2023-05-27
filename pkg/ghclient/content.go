package ghclient

import (
	"encoding/json"
	"fmt"
)

type FileContent struct {
	DownloadUrl string `json:"download_url"`
}

type ContentRequest struct {
	path string
}

func NewContentRequest(user string, repo string, path string) ContentRequest {
	fullPath := fmt.Sprintf("repos/%s/%s/contents/%s", user, repo, path)
	return ContentRequest{fullPath}
}

func (c *GithubClient) GetContent(req ContentRequest) ([]FileContent, error) {
	res, err := c.get(req.path)
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

// If last segment of the provided path has a period, then it is probably a file.
// Also we should check last period index, because hidden files and dirs start with a period.
func (self ContentRequest) isFile() bool {
	p := self.path
	for i := len(p) - 1; i > 0; i-- {
		if p[i] == '/' {
			break
		}
		if p[i] == '.' && p[i-1] != '/' {
			return true
		}
	}
	return false
}
