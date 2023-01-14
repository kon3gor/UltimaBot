package commands

import (
	"dev/kon3gor/ultima/internal/context"
	"dev/kon3gor/ultima/internal/stickers"
)

func ProcessCommand(context *context.Context) {
	command := context.RawUpdate.Message.Command()

	switch command {
	case randomeNameCmd:
		randomName(context)
	case spamCmd:
		spam(context)
	case dailyCmd:
		daily(context)
	case ideaCmd:
		idea(context)
	case chatIdCmd:
		chatId(context)
	case scheduleCmd:
		schedule(context)
	case noteCmd:
		note(context)
	default:
		unknown(context)
	}
}

func unknown(context *context.Context) {
	context.StickerAnswer(stickers.QuestioningAnimeGitl)
	//context.TextAnswer("Sorry, I cannot understand yoy")
}
