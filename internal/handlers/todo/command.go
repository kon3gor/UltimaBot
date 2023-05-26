package todo

import (
	"dev/kon3gor/ultima/internal/appcontext"
	"dev/kon3gor/ultima/internal/ghclient"
	"dev/kon3gor/ultima/internal/guard"
	"fmt"
)

const (
	Cmd      = "todo"
	todoPath = "plans/todoist.md"
)

func ProcessCommand(ctx *appcontext.Context) {
	if err := ctx.Guard(guard.DefaultUserNameGuard); err != nil {
		return
	}

	text := ctx.Args
	content := getTodoContent()

	content = fmt.Sprintf("%s\n- [ ] %s", content, text)
	ghclient.PushContent(ghclient.NewPersonalObsidianRequest(todoPath, content))
	ctx.TextAnswer("Saved!")
}

func getTodoContent() string {
	req := ghclient.NewMyContentRequest("PersonalObsidian", todoPath)
	res, err := ghclient.Default.GetFile(req)
	if err != nil {
		return ""
	}

	return res
}
