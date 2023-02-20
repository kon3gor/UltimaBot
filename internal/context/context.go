package context

import (
	"dev/kon3gor/ultima/internal/fsm"
	"dev/kon3gor/ultima/internal/guard"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Context struct {
	UserName  string
	ChatID    int64
	RawUpdate tgbotapi.Update
	State     *fsm.State

	bot *tgbotapi.BotAPI
}

func CreateFromCommand(update tgbotapi.Update, bot *tgbotapi.BotAPI) *Context {
	chatId := update.Message.Chat.ID
	username := update.Message.From.UserName
	state := fsm.StateStore.GetOrCreateState(chatId)
	return &Context{username, chatId, update, state, bot}
}

func CreateFromCallback(update tgbotapi.Update, bot *tgbotapi.BotAPI) *Context {
	chatId := update.CallbackQuery.Message.Chat.ID
	username := update.CallbackQuery.Message.From.UserName
	state := fsm.StateStore.GetOrCreateState(chatId)
	return &Context{username, chatId, update, state, bot}
}

func (self *Context) CustomAnswer(msg tgbotapi.Chattable) {
	self.sendChattable(msg, false)
}

func (self *Context) TextAnswer(text string) {
	chatId := self.ChatID
	msg := tgbotapi.NewMessage(chatId, text)
	self.sendChattable(msg, false)
}

func (self *Context) MardownAnswer(text string) {
	chatId := self.ChatID
	msg := tgbotapi.NewMessage(chatId, text)
	msg.ParseMode = "MarkdownV2"
	self.sendChattable(msg, false)
}

func (self *Context) Destroyable(msg tgbotapi.Chattable) {
	self.sendChattable(msg, true)
}

func (self *Context) deleteMessageAfter(chatId int64, messageId int, dur time.Duration) {
	msg := tgbotapi.NewDeleteMessage(chatId, messageId)
	timer := time.NewTimer(dur)
	<-timer.C
	self.sendChattable(msg, false)
}

func (self *Context) StickerAnswer(sticker string) {
	chatId := self.ChatID
	msg := tgbotapi.NewSticker(chatId, tgbotapi.FileID(sticker))
	self.sendChattable(msg, false)
}

func (self *Context) sendChattable(msg tgbotapi.Chattable, destroyable bool) {
	message, err := self.bot.Send(msg)
	if err != nil {
		panic(err)
	}
	
	if destroyable {
		go self.deleteMessageAfter(self.ChatID, message.MessageID, 5 * time.Second)
	}
}

func (self *Context) Guard(f guard.GuardFunc) *guard.GuardErr {
	if err := f(self.RawUpdate); err != nil {
		return err
	}

	return nil
}
