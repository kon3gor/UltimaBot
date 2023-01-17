package commands

import (
	"dev/kon3gor/ultima/internal/context"
	"dev/kon3gor/ultima/internal/handlers/daily"
	"dev/kon3gor/ultima/internal/stickers"
	"strings"
)

func ProcessCommand(context *context.Context) {
	command := context.RawUpdate.Message.Command()

	switch command {
	case randomeNameCmd:
		randomName(context)
	case spamCmd:
		spam(context)
	case ideaCmd:
		idea(context)
	case chatIdCmd:
		chatId(context)
	case scheduleCmd:
		schedule(context)
	case noteCmd:
		note(context)
	case daily.Cmd:
		daily.ProcessCommand(context)
	default:
		unknown(context)
	}
}

func getArgs(context *context.Context) []string {
	msg := context.RawUpdate.Message.Text
	parts := strings.Split(msg, " ")
	if len(parts) <= 1 {
		return make([]string, 0)
	} else {
		return parts[1:]
	}
}

func unknown(context *context.Context) {
	context.StickerAnswer(stickers.QuestioningAnimeGitl)
	//context.TextAnswer("Sorry, I cannot understand yoy")
}
