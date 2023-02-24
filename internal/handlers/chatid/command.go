package chatid

import (
	"dev/kon3gor/ultima/internal/appcontext"
	"fmt"
)

const Cmd = "chatId"

func ProcessCommand(context *appcontext.Context) {
	context.TextAnswer(fmt.Sprint(context.ChatID))
}
