package config

import "os"

var Config = struct {
	TgBotToken  string
	VasChatID   string
	MaxChatID   string
	BacePicoApi string
}{
	TgBotToken:  os.Getenv("TG_BOT_TOKEN"),
	VasChatID:   os.Getenv("VAS_CHAT_ID"),
	MaxChatID:   os.Getenv("MAX_CHAT_ID"),
	BacePicoApi: os.Getenv("BACE_PICO_API"),
}
