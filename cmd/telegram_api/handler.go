package telegramapi

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

func handleUpdate(update tgbotapi.Update) (tgbotapi.MessageConfig, error) {
	response := tgbotapi.NewMessage(update.Message.Chat.ID, "hello")
	return response, nil
}
