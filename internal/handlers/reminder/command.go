package reminder

import (
	"dev/kon3gor/ultima/internal/context"
	"dev/kon3gor/ultima/internal/guard"
	"strconv"
	"strings"
	"time"
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
		context.TextAnswer("Not yet available")
	case "new":
		createNewTimer(context, args)
	}

}

func createNewTimer(context *context.Context, args []string) {
	parts := strings.SplitN(args[1], " ", 2)
	duration := timeFromString(parts[0])
	go ticker(context, parts[1], duration)
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

func ticker(context *context.Context, textToSpam string, dur time.Duration) {
	timer := time.NewTimer(dur)
	<-timer.C
	context.TextAnswer(textToSpam)
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
