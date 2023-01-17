package daily

import (
	"dev/kon3gor/ultima/internal/context"
	"fmt"
	"strconv"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const Callback = "daily"

func ProcessCallback(context *context.Context, data string) {
	entries := strings.Split(data, ",")
	if entries[0] == "navigate" {
		total, lower, err := parseNavigateArgs(entries)
		if err != nil {
			panic(err)
		}

		keyBoard := createKeyBoardWithLowerBound(total, lower)
		msgId := context.RawUpdate.CallbackQuery.Message.MessageID
		msg := tgbotapi.NewEditMessageReplyMarkup(context.ChatID, msgId, keyBoard)
		context.CustomAnswer(msg)
	} else if entries[0] == "check" {
		//todo: mark task as completed smth here
		daily := makeGithubRequest()
		indicies := dailiyAsIndList(daily)
		index, err := parseCheckArgs(entries)
		if err != nil {
			panic(err)
		}
		var upper int
		if index+1 == len(indicies) {
			upper = len(daily) - 1
		} else {
			upper = indicies[index+1]
		}
		task := daily[indicies[index]:upper]
		task = strings.ReplaceAll(task, "[ ]", "[x]")
		daily = daily[:indicies[index]] + task + daily[upper:]
		context.TextAnswer(fmt.Sprintf("New daily:\n%s", daily))
	}
}

func parseNavigateArgs(args []string) (int, int, error) {
	if len(args) != 3 {
		//todo: genrate error here
		return -1, -1, nil
	}

	total, err := strconv.Atoi(args[1])
	if err != nil {
		return -1, -1, err
	}
	lower, err := strconv.Atoi(args[2])
	if err != nil {
		return -1, -1, err
	}

	return total, lower, nil
}

func parseCheckArgs(args []string) (int, error) {
	if len(args) != 2 {
		//todo: genrate error here
		return -1, nil
	}

	index, err := strconv.Atoi(args[1])
	if err != nil {
		return -1, err
	}

	return index, nil
}

func createNextKeyboard(total, lower int) tgbotapi.InlineKeyboardMarkup {
	upper := lower + 5
	if upper > total {
		upper = total
	}
	length := upper - lower
	row := make([]tgbotapi.InlineKeyboardButton, length)
	for i := 0; i < length; i++ {
		text := fmt.Sprint(lower + i + 1)
		callback := fmt.Sprintf("daily:check,%d", i)
		row[i] = tgbotapi.NewInlineKeyboardButtonData(text, callback)
	}
	navbar := make([]tgbotapi.InlineKeyboardButton, 0)
	if lower-5 >= 0 {
		prev_data := fmt.Sprintf("daily:navigate,%d,%d", total, lower-5)
		navbar = append(navbar, tgbotapi.NewInlineKeyboardButtonData("<<", prev_data))
	}
	if upper != total {
		next_data := fmt.Sprintf("daily:navigate,%d,%d", total, upper)
		navbar = append(navbar, tgbotapi.NewInlineKeyboardButtonData(">>", next_data))
	}
	return tgbotapi.NewInlineKeyboardMarkup(row, navbar)
}
