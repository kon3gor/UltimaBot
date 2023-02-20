package reminder

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const (
	createButtonText     string = "Create a reminder ⏰"
	litstButtonText             = "List reminders 🗒"
	backButtonText              = "Back"
	periodicButtonText          = "Periodic"
	singleTimeButtonText        = "Single time"
)

func createOrListKeyboard() tgbotapi.InlineKeyboardMarkup {
	return tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(createButtonText, "remind:navigate:create"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(litstButtonText, "remind:navigate:list"),
		),
	)
}

func reminderTypeKeyboard() tgbotapi.InlineKeyboardMarkup {
	return tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(periodicButtonText, "remind:navigate:home"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(singleTimeButtonText, "remind:navigate:home"),
		),
	)
}

func newKeyboardWithBackButtonVararg(from string, rows ...[]tgbotapi.InlineKeyboardButton) tgbotapi.InlineKeyboardMarkup {
	return newKeyboardWithBackButton(from, rows)
}

func newKeyboardWithBackButton(from string, rows [][]tgbotapi.InlineKeyboardButton) tgbotapi.InlineKeyboardMarkup {
	backButton := tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData(backButtonText, fmt.Sprintf("remind:navigate:%s", from)),
	)
	rows = append(rows, backButton)
	return tgbotapi.InlineKeyboardMarkup{InlineKeyboard: rows}
}
