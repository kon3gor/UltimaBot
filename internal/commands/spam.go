package commands

import (
	"dev/kon3gor/ultima/internal/context"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const spamCmd = "spam"

var spammers = make(map[string]chan int8)

func spam(context *context.Context) {
	t := context.RawUpdate.Message.Text
	m := strings.SplitN(t, " ", 2)
	textToSpam := m[1]

	if val, ok := spammers[textToSpam]; ok {
		val <- 0
		delete(spammers, textToSpam)
	} else {
		go ticker(context, textToSpam)
	}
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
