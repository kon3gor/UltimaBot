package callbacks

import (
	"dev/kon3gor/ultima/internal/context"
	"dev/kon3gor/ultima/internal/handlers/daily"
	"strings"
)

func ProcessCallback(context *context.Context) {
	data := context.RawUpdate.CallbackQuery.Data
	callback, args, _ := strings.Cut(data, ":")
	switch callback {
	case randomNameCallback:
		randomName(args, context)
	case daily.Callback:
		daily.ProcessCallback(context, args)
	default:
		unknown(context)
	}
}

func unknown(context *context.Context) {
	context.TextAnswer("Sorry, I cannot understand yoy")
}
