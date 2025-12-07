package bot

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func (s *srv) sendTextMessage(chatID int64, text string) error {
	msg := tgbotapi.NewMessage(chatID, text)
	_, err := s.bot.Send(msg)
	return err
}

func (s *srv) sendPhoto(chatID int64, photo []byte, fileName string) error {
	msg := tgbotapi.NewPhoto(
		chatID,
		tgbotapi.FileBytes{
			Name:  fileName,
			Bytes: photo,
		},
	)
	_, err := s.bot.Send(msg)
	return err
}
