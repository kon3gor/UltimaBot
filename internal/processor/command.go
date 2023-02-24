package processor

import (
	"dev/kon3gor/ultima/internal/appcontext"
	"dev/kon3gor/ultima/internal/handlers/chatid"
	"dev/kon3gor/ultima/internal/handlers/daily"
	"dev/kon3gor/ultima/internal/handlers/idea"
	"dev/kon3gor/ultima/internal/handlers/name"
	"dev/kon3gor/ultima/internal/handlers/note"
	"dev/kon3gor/ultima/internal/handlers/pokemon"
	"dev/kon3gor/ultima/internal/handlers/reminder"
	"dev/kon3gor/ultima/internal/handlers/schedule"
	"dev/kon3gor/ultima/internal/handlers/spam"
	"dev/kon3gor/ultima/internal/stickers"
	"strings"
)

func ProcessCommand(context *appcontext.Context) {
	command := context.RawUpdate.Message.Command()

	switch command {
	case name.Cmd:
		name.ProcessCommand(context)
	case spam.Cmd:
		spam.ProcessCommand(context)
	case idea.Cmd:
		idea.ProcessCommand(context)
	case chatid.Cmd:
		chatid.ProcessCommand(context)
	case schedule.Cmd:
		schedule.ProcessCommand(context)
	case note.Cmd:
		note.ProcessCommand(context)
	case daily.Cmd:
		daily.ProcessCommand(context)
	case pokemon.Cmd:
		pokemon.ProcessCommand(context)
	case reminder.Cmd:
		reminder.ProcessCommand(context)
	default:
		unknownCommand(context)
	}
}

func getArgs(context *appcontext.Context) []string {
	msg := context.RawUpdate.Message.Text
	parts := strings.Split(msg, " ")
	if len(parts) <= 1 {
		return make([]string, 0)
	} else {
		return parts[1:]
	}
}

func unknownCommand(context *appcontext.Context) {
	context.StickerAnswer(stickers.QuestioningAnimeGitl)
}
