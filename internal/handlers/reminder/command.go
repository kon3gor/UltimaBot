package reminder

import (
	"dev/kon3gor/ultima/internal/context"
	"dev/kon3gor/ultima/internal/guard"
	"fmt"
	"math/rand"
	"strconv"
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
	context.TextAnswer(msg)
	//todo: make context start flow and put state into state store
	context.State.StartFlow(Cmd)
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

	return 0 * time.Second
}

func ticker(context *context.Context, textToSpam string) {
	ticker := time.NewTicker(5 * time.Second)
	quit := make(chan int8)
	id := fmt.Sprint(rand.Intn(100))
	context.TextAnswer(id)
	spammers[id] = quit
	for {
		select {
		case <-ticker.C:
			context.TextAnswer(textToSpam)
		case <-quit:
			context.TextAnswer("spam ended")
			ticker.Stop()
			return
		}
	}
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
