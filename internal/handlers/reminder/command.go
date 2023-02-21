package reminder

import (
	"dev/kon3gor/ultima/internal/context"
	"dev/kon3gor/ultima/internal/guard"
	"fmt"
	"math/rand"
	"strconv"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const Cmd = "remind"

var spammers = make(map[string]chan int8)

func ProcessCommand(context *context.Context) {
	if err := context.Guard(guard.DefaultUserNameGuard); err != nil {
		context.TextAnswer(err.Msg)
		return
	}
	guarded(context)
}

func guarded(context *context.Context) {
	msg := tgbotapi.NewMessage(context.ChatID, "What would u like to do?")
	msg.ReplyMarkup = createOrListKeyboard()
	context.CustomAnswer(msg)
}

var reminders map[int]string = make(map[int]string, 0)

func listReminders() tgbotapi.InlineKeyboardMarkup {
	var rows [][]tgbotapi.InlineKeyboardButton
	for id, v := range reminders {
		data := fmt.Sprintf("remind:delete:%d", id)
		button := tgbotapi.NewInlineKeyboardButtonData(v, data)
		rows = append(rows, []tgbotapi.InlineKeyboardButton{button})
	}
	return newKeyboardWithBackButton("home", rows)
}

func timeFromString(strTime string) time.Duration {
	unit := string(strTime[len(strTime)-1])
	num, err := strconv.Atoi(string(strTime[:len(strTime)-1]))
	if err != nil {
		return 0 * time.Second
	}

	switch unit {
	case "s":
		return time.Duration(num) * time.Second
	case "m":
		return time.Duration(num) * time.Minute
	case "h":
		return time.Duration(num) * time.Hour
	}

	return time.Duration(0)
}

func timer(context *context.Context, textToSpam string, dur time.Duration, id int) {
	timer := time.NewTimer(dur)
	<-timer.C
	context.TextAnswer(textToSpam)
	delete(reminders, id)
}

func ticker(context *context.Context, text string, dur time.Duration, id int) {
	ticker := time.NewTicker(dur)
	for range ticker.C {
		if _, ok := reminders[id]; ok {
			context.TextAnswer(text)
		} else {
			ticker.Stop()
		}
	}
}

var reminderDuration string
var reminderMessage string

func ProcessFlow(context *context.Context) {
	switch context.State.CurrentStep() {
	case 0:
		reminderDuration = context.RawUpdate.Message.Text
		context.TextAnswer("Enter reminder message")
		context.State.Next()
	case 1:
		reminderMessage = context.RawUpdate.Message.Text
		context.TextAnswer(fmt.Sprintf("%s %s %s", reminderType, reminderDuration, reminderMessage))
		createReminder(context)
		context.State.FinishFlow()
	}
}

func createReminder(context *context.Context) {
	duration := timeFromString(reminderDuration)
	id := rand.Intn(1000)
	reminders[id] = reminderMessage
	if reminderType == "periodic" {
		go ticker(context, reminderMessage, duration, id)
	} else {
		go timer(context, reminderMessage, duration, id)
	}
}
