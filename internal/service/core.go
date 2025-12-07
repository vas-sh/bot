package service

import "strconv"

type tgboter interface {
	ManageActivity(chatID int64, photo []byte, fileName string) error
}

type srv struct {
	bot       tgboter
	vasChatID int
}

func New(bot tgboter, vasChatIDStr string) (*srv, error) {
	vasChatID, err := strconv.Atoi(vasChatIDStr)
	if err != nil {
		return nil, err
	}
	return &srv{
		bot:       bot,
		vasChatID: vasChatID,
	}, nil
}
