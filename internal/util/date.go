package util

import (
	"fmt"
	"time"
)

func GetCurrentDateAsMillis() (int64, error) {
	t, err := getCurrentTime()
	if err != nil {
		return -1, err
	}
	date := t.UnixMilli() / 86400000
	return date, nil
}

const oneDay time.Duration = 24 * time.Hour

func GetDateAsString(shift int) (string, error) {
	t, err := getCurrentTime()
	if err != nil {
		return "", err
	}
	t.Add(oneDay * time.Duration(shift))
	year, month, day := t.Date()
	return fmt.Sprintf("%d-%02d-%02d", year, int(month), day), nil
}

func getCurrentTime() (time.Time, error) {
	tz, err := time.LoadLocation("Europe/Moscow")
	if err != nil {
		return time.Time{}, err
	}
	return time.Now().In(tz), nil
}
