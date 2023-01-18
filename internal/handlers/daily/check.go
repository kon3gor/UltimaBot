package daily

import (
	"dev/kon3gor/ultima/internal/context"
	"dev/kon3gor/ultima/internal/ghclient"
	"fmt"
	"net/http"
	"regexp"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func checkDaily(context *context.Context, at int) {
	dailyFromMsg := context.RawUpdate.CallbackQuery.Message.Text
	daily := defomratDaily(dailyFromMsg)
	markedDaily := markFinishedTask(daily, at)
	formatted := formatDaily(markedDaily)
	msgId := context.RawUpdate.CallbackQuery.Message.MessageID
	msg := tgbotapi.NewEditMessageText(context.ChatID, msgId, formatted)
	msg.ParseMode = "MarkdownV2"
	msg.ReplyMarkup = context.RawUpdate.CallbackQuery.Message.ReplyMarkup

	if err := pushChangesToGithub(markedDaily); err != nil {
		context.TextAnswer(fmt.Sprint(err))
	} else {
		context.CustomAnswer(msg)
	}
}

func markFinishedTask(daily string, at int) string {
	indicies := dailiyAsIndList(daily)
	var upper int
	if at+1 == len(indicies) {
		upper = len(daily) - 1
	} else {
		upper = indicies[at+1]
	}
	task := daily[indicies[at]:upper]
	task = strings.ReplaceAll(task, "[ ]", "[x]")
	return daily[:indicies[at]] + task + daily[upper:]
}

func pushChangesToGithub(newDaily string) error {
	currentDate, err := getCurrentDate()
	if err != nil {
		return err
	}
	client := &http.Client{}
	path := fmt.Sprintf("plans/daily/%s.md", currentDate)
	req := ghclient.NewPersonalObsidianRequest(path, newDaily)
	ghclient.PushContent(client, req)
	return nil
}

var numRe *regexp.Regexp = regexp.MustCompile(`\d\. `)

func defomratDaily(daily string) string {
	daily = strings.ReplaceAll(daily, "✅", "- [x]")
	daily = strings.ReplaceAll(daily, "❌", "- [ ]")
	daily = strings.ReplaceAll(daily, "    ", "\t")
	daily = numRe.ReplaceAllString(daily, "")

	return daily
}
