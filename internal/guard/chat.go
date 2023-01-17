package guard

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

const (
	myChat = 543162623
)

func MyChat(update tgbotapi.Update) *GuardErr {
	chatId := update.Message.Chat.ID
	if chatId != myChat {
		return NewGuardErr("Sorry, u cannot use it in this chat")
	}

	return nil
}
