package ghclient

import "os"

var Default = defaultClient(os.Getenv("GITHUB_TOKEN"))

func NewMyContentRequest(repo string, path string) ContentRequest {
	return NewContentRequest("kon3gor", repo, path)
}

func GetContent(req ContentRequest) ([]FileContent, error) {
	return Default.GetContent(req)
}

func NewPersonalObsidianRequest(path string, content string) PushRequest {
	return NewPushRequest("kon3gor", "PersonalObsidian", "main", path, content)
}

func PushContent(req PushRequest) error {
	return Default.PushContent(req)
}

func GetFile(req ContentRequest) (string, error) {
	return Default.GetFile(req)
}
