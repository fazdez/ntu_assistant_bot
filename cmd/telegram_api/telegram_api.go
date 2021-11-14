package telegramapi

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type Config struct {
	Token        string
	UpdateConfig tgbotapi.UpdateConfig
}

type TelegramAPIWrapper struct {
	api     *tgbotapi.BotAPI
	config  Config
	updates tgbotapi.UpdatesChannel
}

func New(config Config) (*TelegramAPIWrapper, error) {
	bot, err := tgbotapi.NewBotAPI(config.Token)
	if err != nil {
		return nil, err
	}
	updates, err := bot.GetUpdatesChan(config.UpdateConfig)
	if err != nil {
		return nil, err
	}
	return &TelegramAPIWrapper{bot, config, updates}, nil
}

func (tg *TelegramAPIWrapper) Run() error {
	for update := range tg.updates {
		if err := validateUpdate(update); err != nil {
			return err
		}

		msg, err := handleUpdate(update)
		if err != nil {
			return err
		}

		tg.api.Send(msg)
	}

	return nil
}

func (tg *TelegramAPIWrapper) Close() error {
	tg.api.StopReceivingUpdates()
	return nil
}
