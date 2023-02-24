package reminder

import (
	"dev/kon3gor/ultima/internal/appcontext"
	"strconv"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const Callback = "remind"

func ProcessCallback(context *appcontext.Context, args string) {
	splittedArgs := strings.Split(args, ":")
	command := splittedArgs[0]
	switch command {
	case "navigate":
		navigate(context, splittedArgs[1:])
	case "create":
		create(context, splittedArgs[1:])
	case "delete":
		deleteReminder(context, splittedArgs[1:])
	}

}

var reminderType string

func create(context *appcontext.Context, args []string) {
	reminderType = args[0]
	changeText(context, "Enter duration")
	context.State.StartFlow(Cmd)
}

// todo: check if reminder stil exists
func deleteReminder(context *appcontext.Context, args []string) {
	id, _ := strconv.Atoi(args[0])
	delete(reminders, id)
	context.TextAnswer("Deleted!")
	changeKeyboard(context, listReminders())
}

func navigate(context *appcontext.Context, args []string) {
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

func changeKeyboardAndText(context *appcontext.Context, text string, keyboard tgbotapi.InlineKeyboardMarkup) {
	msgId := context.RawUpdate.CallbackQuery.Message.MessageID
	msg := tgbotapi.NewEditMessageTextAndMarkup(context.ChatID, msgId, text, keyboard)
	context.CustomAnswer(msg)
}

func changeKeyboard(context *appcontext.Context, keyboard tgbotapi.InlineKeyboardMarkup) {
	msgId := context.RawUpdate.CallbackQuery.Message.MessageID
	msg := tgbotapi.NewEditMessageReplyMarkup(context.ChatID, msgId, keyboard)
	context.CustomAnswer(msg)
}

func changeText(context *appcontext.Context, text string) {
	msgId := context.RawUpdate.CallbackQuery.Message.MessageID
	textMsg := tgbotapi.NewEditMessageText(context.ChatID, msgId, text)
	context.CustomAnswer(textMsg)
}
