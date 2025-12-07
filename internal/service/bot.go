package service

func (s *srv) SendPhotoToBot(photo []byte, fileName string) error {
	err := s.bot.ManageActivity(int64(s.vasChatID), photo, fileName)
	if err != nil {
		return nil
	}
	return s.bot.ManageActivity(int64(s.maxChatID), photo, fileName)
}
