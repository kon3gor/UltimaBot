package guard

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

var defaultGuardedUserNames = []string{"eshendo", "zosuku"}

func DefaultUserNameGuard(update tgbotapi.Update) *GuardErr {
	from := update.Message.From.UserName
	if !contains(defaultGuardedUserNames, from) {
		return NewGuardErr("Sorry, you cannot use this commad")
	}

	return nil
}

func contains(slice []string, str string) bool {
	for _, item := range slice {
		if item == str {
			return true
		}
	}
	return false
}
