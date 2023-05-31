package obsidian

import (
	"dev/kon3gor/ultima/internal/github"
	"fmt"
	"time"
)

const (
	dailyPathTemplate = "plans/daily/%s.md"
)

func TodaysDaily() (string, error) {
	currentDate, err := getCurrentDate()
	if err != nil {
		return "", err
	}
	filePath := fmt.Sprintf(dailyPathTemplate, currentDate)
	return github.GetObsidianFile(filePath)
}

func ModifyTodaysDaily(content string) error {
	currentDate, err := getCurrentDate()
	if err != nil {
		return err
	}
	path := fmt.Sprintf(dailyPathTemplate, currentDate)
	return github.SaveObisdianFile(path, content)
}

func GetIdeaUrls() ([]string, error) {
	raw, err := github.GetObsidianFolder("ideas")
	if err != nil {
		return nil, err
	}
	ideas := make([]string, 0, len(raw))
	for _, rawIdea := range raw {
		if rawIdea.IsFile() {
			ideas = append(ideas, rawIdea.DownloadUrl)
		} else {
			//todo: do drill down here
		}
	}

	return ideas, nil
}

func getCurrentDate() (string, error) {
	tz, err := time.LoadLocation("Europe/Moscow")
	if err != nil {
		return "", err
	}

	year, month, day := time.Now().In(tz).Date()
	return fmt.Sprintf("%d-%02d-%02d", year, int(month), day), nil
}
