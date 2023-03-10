package processor

import (
	"dev/kon3gor/ultima/internal/appcontext"
	"dev/kon3gor/ultima/internal/handlers/daily"
	"dev/kon3gor/ultima/internal/handlers/name"
	"dev/kon3gor/ultima/internal/handlers/reminder"
	"strings"
)

func ProcessCallback(context *appcontext.Context) {
	data := context.RawUpdate.CallbackQuery.Data
	callback, args, _ := strings.Cut(data, ":")
	switch callback {
	case name.Callback:
		name.ProcessCallback(context, args)
	case daily.Callback:
		daily.ProcessCallback(context, args)
	case reminder.Callback:
		reminder.ProcessCallback(context, args)
	default:
		unknownCallback(context)
	}
}

func unknownCallback(context *appcontext.Context) {
	context.TextAnswer("Sorry, I cannot understand yoy")
}
