package daily

import (
	"dev/kon3gor/ultima/internal/github"
	"fmt"
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
		realIndex := lower + i
		text := fmt.Sprint(realIndex + 1)
		callback := fmt.Sprintf("daily:check,%d", realIndex)
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

func makeGithubRequest() (string, error) {
	filePathTemplate := "plans/daily/%s.md"
	currentDate, err := getCurrentDate()
	if err != nil {
		return "", err
	}
	filePath := fmt.Sprintf(filePathTemplate, currentDate)
	return github.GetObsidianFile(filePath)
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
