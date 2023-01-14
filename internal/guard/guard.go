package guard

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

type GuardFunc func(tgbotapi.Update) *GuardErr
