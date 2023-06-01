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

func getCurrentDate() (string, error) {
	tz, err := time.LoadLocation("Europe/Moscow")
	if err != nil {
		return "", err
	}

	year, month, day := time.Now().In(tz).Date()
	return fmt.Sprintf("%d-%02d-%02d", year, int(month), day), nil
}
