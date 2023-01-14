package commands

import (
	"dev/kon3gor/ultima/internal/ghclient"
	"dev/kon3gor/ultima/internal/context"
	"dev/kon3gor/ultima/internal/guard"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"strings"
)

const dailyCmd = "daily"

func daily(context *context.Context) {
	if err := context.Guard(guard.DefaultUserNameGuard); err != nil {
		context.TextAnswer(err.Msg)
		return
	}
	dailyGuarded(context)
}

func dailyGuarded(context *context.Context) {
	daily := makeGithubRequest()
	context.TextAnswer(daily)
}

type RequestBody struct {
	DownloadUrl string `json:"download_url"`
}

func makeGithubRequest() string {
	filePathTemplate := "plans/daily/%s.md"
	currentDate, err := getCurrentDate()
	if err != nil {
		return fmt.Sprintf("Error while getting current date: %s", err)
	}
	filePath := fmt.Sprintf(filePathTemplate, currentDate)
	client := &http.Client{}
	req := ghclient.NewMyContentRequest("PersonalObsidian", filePath)
	content, _ := ghclient.GetContent(client, req)

	res, err := client.Get(content[0].DownloadUrl)
	if err != nil {
		return fmt.Sprintf("Error while fetching donwload: %s", err)
	}

	defer res.Body.Close()

	bodyBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return fmt.Sprintf("error parsing bytes: %s", err)
	}

	return strings.ReplaceAll(string(bodyBytes), "\t", "    ")
}

func getCurrentDate() (string, error) {
	tz, err := time.LoadLocation("Europe/Moscow")
	if err != nil {
		return "", err
	}

	year, month, day := time.Now().In(tz).Date()
	return fmt.Sprintf("%d-%02d-%02d", year, int(month), day), nil
}
