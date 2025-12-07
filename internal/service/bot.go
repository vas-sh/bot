package service

func (s *srv) SendPhotoToBot(photo []byte, fileName string) error {
	return s.bot.ManageActivity(int64(s.vasChatID), photo, fileName)
}
