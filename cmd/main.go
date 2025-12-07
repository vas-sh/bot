package main

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/vas-sh/bot/internal/bot"
	"github.com/vas-sh/bot/internal/config"
	"github.com/vas-sh/bot/internal/handlers"
	"github.com/vas-sh/bot/internal/handlers/bothandlers"
	"github.com/vas-sh/bot/internal/service"
)

func main() {
	cfg := config.Config

	tgBotAPI, err := tgbotapi.NewBotAPI(cfg.TgBotToken)
	if err != nil {
		panic(err)
	}
	tgBot := bot.New(tgBotAPI)
	go tgBot.Updates()

	server := handlers.New()

	router := server.Router()

	botSrv, err := service.New(tgBot, cfg.VasChatID)
	if err != nil {
		panic(err)
	}

	bothandlers.New(botSrv).Register(router)

	log.Println("Starting server...")
	err = server.Run()
	if err != nil {
		panic(err)
	}
}
