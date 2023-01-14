package commands

import (
	"dev/kon3gor/ultima/internal/context"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const randomeNameCmd string = "name"

func randomName(context *context.Context) {
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
