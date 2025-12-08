package bot

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/vas-sh/bot/internal/models"
)

func (s *srv) Updates() {
	s.createMenu()
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := s.bot.GetUpdatesChan(u)
	for update := range updates {
		go s.manageUpdate(update)
	}
}

func (s *srv) createMenu() {
	commands := []tgbotapi.BotCommand{
		{
			Command:     string(models.TakePhoto),
			Description: "Take photo",
		},
	}
	cfg := tgbotapi.NewSetMyCommands(commands...)
	_, err := s.bot.Request(cfg)
	if err != nil {
		log.Println(err.Error())
	}
}

func (s *srv) manageUpdate(update tgbotapi.Update) {
	if update.Message != nil {
		err := s.handleTextMessage(update.Message)
		if err != nil {
			log.Println(err.Error())
		}
		return
	}
	if update.CallbackQuery != nil {
		callback := tgbotapi.NewCallback(update.CallbackQuery.ID, "")
		_, err := s.bot.Request(callback)
		if err != nil {
			log.Println(err.Error())
		}
		err = s.handleCallback(update.CallbackQuery)
		if err != nil {
			log.Println(err.Error())
		}
	}
}
