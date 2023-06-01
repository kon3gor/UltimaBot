package todo

import (
	"dev/kon3gor/ultima/internal/appcontext"
	"dev/kon3gor/ultima/internal/service/github"
	"dev/kon3gor/ultima/internal/guard"
	"fmt"
	"strings"
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

	content = fmt.Sprintf("%s\n- [ ] %s", strings.Trim(content, "\n"), text)
	github.SaveObisdianFile(todoPath, content)
	ctx.TextAnswer("Saved!")
}

func getTodoContent() string {
	todo, err := github.GetObsidianFile(todoPath)
	if err != nil {
		return ""
	}

	return todo
}
