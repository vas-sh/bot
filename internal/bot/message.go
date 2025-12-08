package bot

import (
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/vas-sh/bot/internal/models"
)

func (s *srv) handleTextMessage(message *tgbotapi.Message) error {
	text := message.Text
	chatID := message.Chat.ID
	if strings.HasPrefix(text, string(models.StartCommand)) {
		return s.sendTextMessage(chatID,
			`
			Welcome to Pico Bot!
			Here, you will receive messages if our system detects activity.`,
		)
	}
	if text == string(models.TakePhoto) {
		err := s.client.TakePhoto()
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *srv) ManageActivity(chatID int64, photo []byte, fileName string) error {
	err := s.sendTextMessage(chatID, "Detected activity!")
	if err != nil {
		return err
	}
	return s.sendPhoto(chatID, photo, fileName)
}
