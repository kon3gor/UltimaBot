package note

import (
	"dev/kon3gor/ultima/internal/appcontext"
	"dev/kon3gor/ultima/internal/ghclient"
	"dev/kon3gor/ultima/internal/guard"
	"fmt"
	"net/http"
	"strings"
)

const Cmd = "note"
const titleLen = 12

func ProcessCommand(context *appcontext.Context) {
	if err := context.Guard(guard.DefaultUserNameGuard); err != nil {
		context.TextAnswer(err.Msg)
		return
	}
	guarded(context)
}

func guarded(context *appcontext.Context) {
	fullMessage := context.RawUpdate.Message.Text
	_, note, _ := strings.Cut(fullMessage, " ")
	titleLen := titleLen
	if len(note) < titleLen {
		titleLen = len(note)
	}
	title := note[:titleLen-1]
	path := fmt.Sprintf("notes/junk/%s.md", title)
	pushReq := ghclient.NewMyPushRequest("PersonalObsidian", "main", path, note)
	client := &http.Client{}
	if err := ghclient.PushContent(client, pushReq); err != nil {
		context.TextAnswer("Error !!!")
	} else {
		context.TextAnswer("Note saved!")
	}
}
