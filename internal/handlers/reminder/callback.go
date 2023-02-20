package reminder

import (
	"dev/kon3gor/ultima/internal/context"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const Callback = "remind"

func ProcessCallback(context *context.Context, args string) {
	splittedArgs := strings.Split(args, ":")
	command := splittedArgs[0]
	switch command {
	case "navigate":
		navigate(context, splittedArgs[1:])
	}

}

func navigate(context *context.Context, args []string) {
	dest := args[0]
	switch dest {
	case "list":
		changeKeyboardAndText(context, "Current reminders", listReminders())
	case "home":
		changeKeyboardAndText(context, "What would u like to do?", createOrListKeyboard())
	case "create":
		changeKeyboardAndText(context, "Which type of reminder do u want to create?", reminderTypeKeyboard())
	}
}

func changeKeyboardAndText(context *context.Context, text string, keyboard tgbotapi.InlineKeyboardMarkup) {
	msgId := context.RawUpdate.CallbackQuery.Message.MessageID
	msg := tgbotapi.NewEditMessageTextAndMarkup(context.ChatID, msgId, text, keyboard)
	context.CustomAnswer(msg)
}

func changeText(context *context.Context, text string) {
	msgId := context.RawUpdate.CallbackQuery.Message.MessageID
	textMsg := tgbotapi.NewEditMessageText(context.ChatID, msgId, text)
	context.CustomAnswer(textMsg)
}
