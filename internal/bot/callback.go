package bot

import (
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/vas-sh/bot/internal/models"
)

func (s *srv) handleCallback(callback *tgbotapi.CallbackQuery) error {
	data := callback.Data
	chatID := callback.Message.Chat.ID
	if data == string(models.TakePhoto) {
		pathToPhoto, err := s.client.TakePhoto()
		if err != nil {
			return err
		}
		data, err := os.ReadFile(pathToPhoto)
		if err != nil {
			return err
		}
		return s.ManageActivity(chatID, data, pathToPhoto)
	}
	return nil
}
