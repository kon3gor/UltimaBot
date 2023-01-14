package commands

import (
	"dev/kon3gor/ultima/internal/ghclient"
	"dev/kon3gor/ultima/internal/context"
	"dev/kon3gor/ultima/internal/guard"
	"fmt"
	"net/http"
	"strings"
)

const noteCmd = "note"

func note(context *context.Context) {
	if err := context.Guard(guard.DefaultUserNameGuard); err != nil {
		context.TextAnswer(err.Msg)
		return
	}
}

func noteInternal(context *context.Context) {
	fullMessage := context.RawUpdate.Message.Text
	_, note, _ := strings.Cut(fullMessage, " ")
	titleLen := 7
	if len(note) < 7 {
		titleLen = len(note)
	}
	title := note[:titleLen-1]
	path := fmt.Sprintf("notes/%s.md", title)
	pushReq := ghclient.NewMyPushRequest("PersonalObsidian", "main", path, note)
	client := &http.Client{}
	if err := ghclient.PushContent(client, pushReq); err != nil {
		context.TextAnswer("Error !!!")
	} else {
		context.TextAnswer("Note saved!")
	}
}
