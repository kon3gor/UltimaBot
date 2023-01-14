package callbacks

import (
	"dev/kon3gor/ultima/internal/context"
	"strings"
)

func ProcessCallback(context *context.Context) {
	data := context.RawUpdate.CallbackQuery.Data
	splitted := strings.SplitN(data, ":", 2)
	switch splitted[0] {
	case randomNameCallback:
		randomName(splitted[1], context)
	default:
		unknown(context)
	}
}

func unknown(context *context.Context) {
	context.TextAnswer("Sorry, I cannot understand yoy")
}
