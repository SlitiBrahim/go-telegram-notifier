package config

import (
	"fmt"
	"os"
)

// TODO: better solution ?
var Config = make(map[string]interface{})

func init() {
	Config["APP_PORT"] = os.Getenv("APP_PORT")
	Config["TG_CHAT_ID"] = os.Getenv("TG_CHAT_ID")
	Config["TG_BOT_TOKEN"] = os.Getenv("TG_BOT_TOKEN")
	Config["TG_API_BOT_BASE_URL"] = fmt.Sprintf("https://api.telegram.org/bot%s/", Config["TG_BOT_TOKEN"])
}
