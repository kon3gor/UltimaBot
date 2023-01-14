package main

import "dev/kon3gor/ultima/internal/telegram"

func main() {
	if err := telegram.InitTelegramBot(); err != nil {
		panic(err)
	}
	telegram.StartPolling()
}
