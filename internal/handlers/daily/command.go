package daily

import (
	"dev/kon3gor/ultima/internal/appcontext"
	"dev/kon3gor/ultima/internal/guard"
	"dev/kon3gor/ultima/internal/util"
	"fmt"
	"log"

	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const (
	Cmd     = "daily"
	eshendo = "eshendo"
	zosuku  = "zosuku"
)

func ProcessCommand(context *appcontext.Context) {
	if err := context.Guard(guard.DefaultUserNameGuard); err != nil {
		context.TextAnswer(err.Msg)
		return
	}
	dailyGuarded(context)
}

func dailyGuarded(context *appcontext.Context) {
	args := strings.Split(context.Args, " ")
	username := context.UserName
	if len(args) > 0 && args[0] != "" {
		username = args[0]
	}
	if username == eshendo {
		mineDaily(context)
	} else {
		sunshineDaily(context)
	}
}

func mineDaily(context *appcontext.Context) {
	raw_daily, err := makeGithubRequest()
	if err != nil {
		log.Println(err)
		context.TextAnswer("Smth went wrong")
		return
	}
	count := len(dailiyAsIndList(raw_daily))
	daily := util.EscapeFakeMarkdown(formatDaily(raw_daily))
	msg := tgbotapi.NewMessage(context.ChatID, daily)
	msg.ParseMode = "MarkdownV2"
	if context.UserName == eshendo {
		msg.ReplyMarkup = createKeyBoardWithLowerBound(count, 0)
	}
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
