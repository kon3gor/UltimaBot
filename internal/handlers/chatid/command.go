package chatid

import (
	"dev/kon3gor/ultima/internal/context"
	"fmt"
)

const Cmd = "chatId"

func ProcessCommand(context *context.Context) {
	context.TextAnswer(fmt.Sprint(context.ChatID))
}
