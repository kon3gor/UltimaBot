package messages

import (
	"dev/kon3gor/ultima/internal/context"
	"dev/kon3gor/ultima/internal/guard"
)

func Process(context *context.Context) {
	if err := context.Guard(guard.DefaultUserNameGuard); err != nil {
		return 
	}

	if context.RawUpdate.Message.Sticker != nil {
		context.TextAnswer(context.RawUpdate.Message.Sticker.FileID)
	}
}
