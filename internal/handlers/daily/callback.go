package daily

import (
	"dev/kon3gor/ultima/internal/context"
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
		index, err := parseCheckArgs(entries)
		if err != nil {
			panic(err)
		}
		checkDaily(context, index)
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
