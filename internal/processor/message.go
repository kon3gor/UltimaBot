package processor

import (
	"dev/kon3gor/ultima/internal/appcontext"
	"dev/kon3gor/ultima/internal/guard"
	"dev/kon3gor/ultima/internal/stickers"
)

func Process(context *appcontext.Context) {
	if context.RawUpdate.Message.Sticker != nil {
		processSticker(context)
	}
}
func processSticker(context *appcontext.Context) {
	if err := context.Guard(guard.DefaultUserNameGuard); err != nil {
		return
	}
	fileID := context.RawUpdate.Message.Sticker.FileID
	fileUniqueID := context.RawUpdate.Message.Sticker.FileUniqueID
	if err := stickers.SaveSticker(fileUniqueID, fileID); err != nil {
		panic(err)
	}
}
