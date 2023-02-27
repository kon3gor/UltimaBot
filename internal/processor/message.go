package processor

import (
	"dev/kon3gor/ultima/internal/appcontext"
	"dev/kon3gor/ultima/internal/guard"
	"dev/kon3gor/ultima/internal/handlers/reminder"
)

func Process(context *appcontext.Context) {
	if context.RawUpdate.Message.Sticker != nil {
		processSticker(context)
	}

	switch context.State.CurrentCmd() {
	case reminder.Cmd:
		reminder.ProcessFlow(context)
	}
}

func processSticker(context *appcontext.Context) {
	if err := context.Guard(guard.DefaultUserNameGuard); err != nil {
		return
	}
	context.TextAnswer(context.RawUpdate.Message.Sticker.FileID)
}
