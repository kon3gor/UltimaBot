package telegram

import (
	"dev/kon3gor/ultima/internal/appcontext"
	"dev/kon3gor/ultima/internal/processor"
	"dev/kon3gor/ultima/internal/stickers"
	"fmt"
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
	bot.Debug = true

	for update := range updates {
		ProcessUpdate(update)
	}
}

func ProcessUpdate(update tgbotapi.Update) {
	if update.Message != nil && update.Message.IsCommand() {
		context := appcontext.CreateFromCommand(update, bot)
		processor.ProcessCommand(context)
	} else if update.CallbackQuery != nil {
		context := appcontext.CreateFromCallback(update, bot)
		processCallback(context)
	} else if update.Message != nil {
		context := appcontext.CreateFromCommand(update, bot)
		processor.Process(context)
	} else if update.InlineQuery != nil {
		context := appcontext.CreateFromInlineQuery(update, bot)
		processInlineQuery(context)
	}
}

func processCallback(context *appcontext.Context) {
	callback := tgbotapi.NewCallback(context.RawUpdate.CallbackQuery.ID, context.RawUpdate.CallbackQuery.Data)
	if _, err := bot.Request(callback); err != nil {
		panic(err)
	}
	processor.ProcessCallback(context)
}

func processInlineQuery(ctx *appcontext.Context) {
	id := ctx.RawUpdate.InlineQuery.ID
	results := make([]interface{}, 0)
	stickers, err := stickers.GetStickers()
	if err != nil {
		return
	}
	for i, sticker := range stickers {
		inlineSticker := tgbotapi.NewInlineQueryResultCachedSticker(
			fmt.Sprint(i),
			sticker,
			"",
		)
		results = append(results, inlineSticker)
	}
	conf := tgbotapi.InlineConfig{
		InlineQueryID: id,
		IsPersonal:    true,
		CacheTime:     0,
		Results:       results,
	}
	if _, err := bot.Request(conf); err != nil {
		panic(err)
	}
}
