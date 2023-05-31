package daily

import (
	"fmt"
	"regexp"

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

var re *regexp.Regexp = regexp.MustCompile(`\t*- \[(x| )\]`)

func dailiyAsIndList(daily string) []int {
	indicies := re.FindAllStringIndex(daily, -1)
	var res []int
	for _, s := range indicies {
		res = append(res, s[0])
	}
	return res
}
