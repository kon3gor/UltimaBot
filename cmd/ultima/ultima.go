package main

import (
	"dev/kon3gor/ultima/internal/telegram"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	if err := telegram.InitTelegramBot(); err != nil {
		panic(err)
	}

	startServer()
}

func startServer() {
	http.HandleFunc("/handle", handler)
	http.HandleFunc("/ping", ping)

	log.Fatalln(http.ListenAndServe(fmt.Sprintf("0.0.0.0:%s", "80"), nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	var res tgbotapi.Update
	if err := json.NewDecoder(r.Body).Decode(&res); err != nil {
		panic(err)
	}

	telegram.ProcessUpdate(res)
	fmt.Fprintf(w, "ok")
}

func ping(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "pong")
}
