package daily

import (
	"dev/kon3gor/ultima/internal/context"
	"dev/kon3gor/ultima/internal/guard"
	"fmt"
	"log"

	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const Cmd = "daily"

func ProcessCommand(context *context.Context) {
	if err := context.Guard(guard.DefaultUserNameGuard); err != nil {
		context.TextAnswer(err.Msg)
		return
	}
	dailyGuarded(context)
}

func dailyGuarded(context *context.Context) {
	raw_daily, err := makeGithubRequest()
	if err != nil {
		log.Println(err)
		context.TextAnswer("Smth went wrong")
		return 
	}
	count := len(dailiyAsIndList(raw_daily))
	daily := formatDaily(raw_daily)
	msg := tgbotapi.NewMessage(context.ChatID, daily)
	msg.ParseMode = "MarkdownV2"
	msg.ReplyMarkup = createKeyBoardWithLowerBound(count, 0)
	context.CustomAnswer(msg)
}

func formatDaily(daily string) string {
	entries := dailiyAsIndList(daily)
	fmt.Println(entries)
	for i, n := range entries {
		rn := n + 5*i
		daily = fmt.Sprintf("%s`%d.` %s", daily[:rn], i+1, daily[rn:])
	}

	daily = strings.ReplaceAll(daily, "- [ ]", "❌")
	daily = strings.ReplaceAll(daily, "- [x]", "✅")
	daily = strings.ReplaceAll(daily, "-", "\\-")
	daily = strings.ReplaceAll(daily, "\t", "    ")

	fmt.Printf("\n%s\n", daily)

	return daily

}
