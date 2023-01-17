package ghclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type PushRequest struct {
	path    string
	content string
	user    string
	repo    string
	branch  string
}

type ShaHolder struct {
	Sha string `json:"sha"`
}

type RefResponse struct {
	Object ShaHolder `json:"object"`
}

type LastCommit struct {
	Tree ShaHolder `json:"tree"`
}

type Tree struct {
	Path    string `json:"path"`
	T       string `json:"type"`
	Mode    string `json:"mode"`
	Content string `json:"content"`
}

type PostTree struct {
	Owner    string `json:"owner"`
	Repo     string `json:"repo"`
	BaseTree string `json:"base_tree"`
	Tree     []Tree `json:"tree"`
}

type PostCommitRequest struct {
	Parents []string `json:"parents"`
	Tree    string   `json:"tree"`
	Message string   `json:"message"`
}

func NewPushRequest(owner string, repo string, branch string, path string, content string) PushRequest {
	return PushRequest{path, content, owner, repo, branch}
}

func NewMyPushRequest(repo string, branch string, path string, content string) PushRequest {
	return NewPushRequest("kon3gor", repo, branch, path, content)
}

func NewPersonalObsidianRequest(path string, content string) PushRequest {
	return NewPushRequest("kon3gor", "PersonalObsidian", "main", path, content)
}

func PushContent(client *http.Client, req PushRequest) error {
	lastCommitSha, err := getLastCommitSha(client, req)
	if err != nil {
		return err
	}
	treeSha, err := getLastCommitTreeSha(client, req, lastCommitSha)
	if err != nil {
		return err
	}
	newTreeSha, err := postNewTree(client, req, treeSha)
	if err != nil {
		return err
	}
	newCommitSha, err := postCommit(client, req, lastCommitSha, newTreeSha)
	if err != nil {
		return err
	}
	if err := patchHead(client, req, newCommitSha); err != nil {
		return err
	}
	return nil
}

func getLastCommitSha(client *http.Client, req PushRequest) (string, error) {
	headUrl := fmt.Sprintf("repos/%s/%s/git/refs/heads/%s", req.user, req.repo, req.branch)
	response, err := get(client, headUrl)
	defer response.Body.Close()
	if err != nil {
		return "", err
	}

	var body RefResponse
	if err = json.NewDecoder(response.Body).Decode(&body); err != nil {
		return "", err
	}
	return body.Object.Sha, nil
}

func getLastCommitTreeSha(client *http.Client, req PushRequest, sha string) (string, error) {
	headUrl := fmt.Sprintf("repos/%s/%s/git/commits/%s", req.user, req.repo, sha)
	response, err := get(client, headUrl)
	defer response.Body.Close()
	if err != nil {
		return "", err
	}

	var body LastCommit
	if err = json.NewDecoder(response.Body).Decode(&body); err != nil {
		return "", err
	}
	return body.Tree.Sha, nil
}

func (self PostTree) String() string {
	return fmt.Sprintf("base: %s", self.BaseTree)
}

func postNewTree(client *http.Client, req PushRequest, sha string) (string, error) {
	url := fmt.Sprintf("repos/%s/%s/git/trees", req.user, req.repo)
	tree := Tree{req.path, "blob", "100644", req.content}
	body := PostTree{"kon3gor", "PersonalObsidian", sha, []Tree{tree}}
	var buffer io.ReadWriter = new(bytes.Buffer)

	if err := json.NewEncoder(buffer).Encode(body); err != nil {
		return "", err
	}
	response, err := post(client, url, buffer)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()
	var newTreeSha ShaHolder
	if err := json.NewDecoder(response.Body).Decode(&newTreeSha); err != nil {
		return "", err
	}

	return newTreeSha.Sha, nil
}

func postCommit(client *http.Client, req PushRequest, lastCommitSha string, newTreeSha string) (string, error) {
	url := fmt.Sprintf("repos/%s/%s/git/commits", req.user, req.repo)
	body := PostCommitRequest{[]string{lastCommitSha}, newTreeSha, "damn"}
	var buffer io.ReadWriter = new(bytes.Buffer)

	if err := json.NewEncoder(buffer).Encode(body); err != nil {
		return "", err
	}
	response, err := post(client, url, buffer)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()
	var newCommitSha ShaHolder
	if err := json.NewDecoder(response.Body).Decode(&newCommitSha); err != nil {
		return "", err
	}

	return newCommitSha.Sha, nil
}

func patchHead(client *http.Client, req PushRequest, sha string) error {
	headUrl := fmt.Sprintf("repos/%s/%s/git/refs/heads/%s", req.user, req.repo, req.branch)
	body := ShaHolder{sha}
	var buffer io.ReadWriter = new(bytes.Buffer)
	if err := json.NewEncoder(buffer).Encode(body); err != nil {
		return err
	}
	_, err := patch(client, headUrl, buffer)
	if err != nil {
		return err
	}

	return nil
}
