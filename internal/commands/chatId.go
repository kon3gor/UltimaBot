package commands

import (
	"dev/kon3gor/ultima/internal/context"
	"fmt"
)

const chatIdCmd = "chatId"

func chatId(context *context.Context) {
	context.TextAnswer(fmt.Sprint(context.ChatID))
}
