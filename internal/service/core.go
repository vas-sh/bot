package service

import "strconv"

type tgboter interface {
	ManageActivity(chatID int64, photo []byte, fileName string) error
}

type srv struct {
	bot       tgboter
	vasChatID int
	maxChatID int
}

func New(bot tgboter, vasChatIDStr, maxChatIDStr string) (*srv, error) {
	vasChatID, err := strconv.Atoi(vasChatIDStr)
	if err != nil {
		return nil, err
	}
	maxChatID, err := strconv.Atoi(maxChatIDStr)
	if err != nil {
		return nil, err
	}
	return &srv{
		bot:       bot,
		vasChatID: vasChatID,
		maxChatID: maxChatID,
	}, nil
}
