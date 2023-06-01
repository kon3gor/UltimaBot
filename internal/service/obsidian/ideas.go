package obsidian

import (
	"dev/kon3gor/ultima/internal/service/github"
)

func GetIdeaUrls() ([]string, error) {
	ideas := drillFolder("ideas")
	return ideas, nil
}

func drillFolder(folderPath string) []string {
	stack := make([]string, 1)
	stack[0] = folderPath
	files := make([]string, 0)

	for len(stack) > 0 {
		last := len(stack) - 1
		path := stack[last]
		stack = stack[:last]

		raw, err := github.GetObsidianFolder(path)
		if err != nil {
			continue
		}

		for _, entry := range raw {
			if entry.IsFile() {
				files = append(files, entry.DownloadUrl)
			} else {
				stack = append(stack, entry.Path)
			}
		}

	}

	return files
}
