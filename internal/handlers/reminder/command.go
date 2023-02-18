package reminder

import (
	"dev/kon3gor/ultima/internal/context"
	"dev/kon3gor/ultima/internal/guard"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
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
	msg := context.RawUpdate.Message.Text
	args := strings.SplitN(msg, " ", 3)
	subCommand := args[1]
	processSubCommand(context, subCommand, args[1:])
}

func processSubCommand(context *context.Context, subCommand string, args []string) {
	switch subCommand {
	case "list":
		if len(reminders) > 0 {
			msg := tgbotapi.NewMessage(context.ChatID, "Currently running reminders")
			msg.ReplyMarkup = listReminders()
			context.CustomAnswer(msg)
		} else {
			context.TextAnswer("No reminders running")
		}
	case "new":
		createNewTimer(context, args)
	}

}

var reminders map[int]string = make(map[int]string, 0)

func listReminders() tgbotapi.InlineKeyboardMarkup {
	var rows [][]tgbotapi.InlineKeyboardButton
	for id, v := range reminders {
		data := fmt.Sprintf("remind:%d", id)
		button := tgbotapi.NewInlineKeyboardButtonData(v, data)
		rows = append(rows, []tgbotapi.InlineKeyboardButton{button})
	}
	return tgbotapi.InlineKeyboardMarkup{InlineKeyboard: rows}
}

func createNewTimer(context *context.Context, args []string) {
	id := rand.Intn(1000)
	parts := strings.SplitN(args[1], " ", 2)
	duration := timeFromString(parts[0])
	reminders[id] = parts[1][:10]
	go ticker(context, parts[1], duration, id)
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

func ticker(context *context.Context, textToSpam string, dur time.Duration, id int) {
	timer := time.NewTimer(dur)
	<-timer.C
	context.TextAnswer(textToSpam)
	delete(reminders, id)
}

func ProcessFlow(context *context.Context) {
	switch context.State.CurrentStep() {
	case 0:
		context.TextAnswer("jajaja")
		context.State.Next()
	case 1:
		context.TextAnswer("hahahah")
		context.State.FinishFlow()
	}
}
