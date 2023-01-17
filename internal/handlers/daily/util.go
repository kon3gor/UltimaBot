package daily

import (
	"dev/kon3gor/ultima/internal/ghclient"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const buttonsInRow = 5

func createKeyBoardWithLowerBound(total, lower int) tgbotapi.InlineKeyboardMarkup {
	upper := lower + buttonsInRow
	if upper > total {
		upper = total
	}
	length := upper - lower
	row := make([]tgbotapi.InlineKeyboardButton, length)
	for i := 0; i < length; i++ {
		text := fmt.Sprint(lower + i + 1)
		callback := fmt.Sprintf("daily:check,%d,%d", i, total)
		row[i] = tgbotapi.NewInlineKeyboardButtonData(text, callback)
	}
	if total <= 5 {
		return tgbotapi.NewInlineKeyboardMarkup(row)
	}

	navbar := make([]tgbotapi.InlineKeyboardButton, 0)
	if lower-buttonsInRow >= 0 {
		prev_data := fmt.Sprintf("daily:navigate,%d,%d", total, lower-buttonsInRow)
		navbar = append(navbar, tgbotapi.NewInlineKeyboardButtonData("<<", prev_data))
	}
	if upper != total {
		next_data := fmt.Sprintf("daily:navigate,%d,%d", total, upper)
		navbar = append(navbar, tgbotapi.NewInlineKeyboardButtonData(">>", next_data))
	}
	return tgbotapi.NewInlineKeyboardMarkup(row, navbar)
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

	return string(bodyBytes)
}

func getCurrentDate() (string, error) {
	tz, err := time.LoadLocation("Europe/Moscow")
	if err != nil {
		return "", err
	}

	year, month, day := time.Now().In(tz).Date()
	return fmt.Sprintf("%d-%02d-%02d", year, int(month), day), nil
}

var re *regexp.Regexp = regexp.MustCompile(`\t*- \[(x| )\]`)

func dailiyAsIndList(daily string) []int {
	indicies := re.FindAllStringIndex(daily, -1)
	var res []int
	for _, s := range indicies {
		res = append(res, s[0])
	}
	return res
}
