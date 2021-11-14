package main

import (
	telegramapi "github.com/fazdez/ntu_assistant_bot/cmd/telegram_api"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {
	//do something
	t, _ := telegramapi.New(telegramapi.Config{Token: "-", UpdateConfig: tgbotapi.UpdateConfig{}})
	t.Run()
}
