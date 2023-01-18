package processor

import (
	"dev/kon3gor/ultima/internal/context"
	"dev/kon3gor/ultima/internal/handlers/daily"
	"dev/kon3gor/ultima/internal/handlers/name"
	"strings"
)

func ProcessCallback(context *context.Context) {
	data := context.RawUpdate.CallbackQuery.Data
	callback, args, _ := strings.Cut(data, ":")
	switch callback {
	case name.Callback:
		name.ProcessCallback(context, args)
	case daily.Callback:
		daily.ProcessCallback(context, args)
	default:
		unknownCallback(context)
	}
}

func unknownCallback(context *context.Context) {
	context.TextAnswer("Sorry, I cannot understand yoy")
}
