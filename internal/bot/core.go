package bot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type boter interface {
	Request(c tgbotapi.Chattable) (*tgbotapi.APIResponse, error)
	Send(c tgbotapi.Chattable) (tgbotapi.Message, error)
	GetUpdatesChan(config tgbotapi.UpdateConfig) tgbotapi.UpdatesChannel
}

type picoclienter interface {
	TakePhoto() (string, error)
}

type srv struct {
	bot    boter
	client picoclienter
}

func New(bot boter, client picoclienter) *srv {
	return &srv{
		bot:    bot,
		client: client,
	}
}
