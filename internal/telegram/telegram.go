package telegram

import (
	"dev/kon3gor/ultima/internal/callbacks"
	"dev/kon3gor/ultima/internal/commands"
	"dev/kon3gor/ultima/internal/context"
	"dev/kon3gor/ultima/internal/messages"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var bot *tgbotapi.BotAPI

func InitTelegramBot() error {
	token := os.Getenv("TELEGRAM_TOKEN")
	tgbot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return err
	}

	bot = tgbot

	return nil
}

func StartPolling() {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		ProcessUpdate(update)
	}
}

func ProcessUpdate(update tgbotapi.Update) {
	if update.Message != nil && update.Message.IsCommand() {
		context := context.CreateFromCommand(update, bot)
		commands.ProcessCommand(context)
	} else if update.CallbackQuery != nil {
		context := context.CreateFromCallback(update, bot)
		processCallback(context)
	} else if update.Message != nil {
		context := context.CreateFromCommand(update, bot)
		messages.Process(context)
	}
}

func processCallback(context *context.Context) {
	callback := tgbotapi.NewCallback(context.RawUpdate.CallbackQuery.ID, context.RawUpdate.CallbackQuery.Data)
	if _, err := bot.Request(callback); err != nil {
		panic(err)
	}
	callbacks.ProcessCallback(context)
}

