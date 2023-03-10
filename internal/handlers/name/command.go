package name

import (
	"dev/kon3gor/ultima/internal/appcontext"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const Cmd string = "name"

func ProcessCommand(context *appcontext.Context) {
	keyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Male", "name:male"),
			tgbotapi.NewInlineKeyboardButtonData("Female", "name:female"),
		),
	)

	msg := tgbotapi.NewMessage(context.ChatID, "Gender?")
	msg.ReplyMarkup = keyboard
	context.CustomAnswer(msg)
}
