package api

import (
	"bytes"
	"encoding/json"
	"go-telegram-notifier/src/config"
	"go-telegram-notifier/src/helper"
	"net/http"
	"net/url"
	"path"
)

type Message struct {
	ChatID string `json:"chat_id"`
	Text   string `json:"text"`
}

func getSendMessageURL() string {
	baseURL, err := url.Parse(config.Config["TG_API_BOT_BASE_URL"].(string))
	helper.FailOnError(err)

	baseURL.Path = path.Join(baseURL.Path, "/sendMessage")

	return baseURL.String()
}

func sendMessage(message Message) (*http.Response, error) {
	body, err := json.Marshal(message)
	helper.FailOnError(err)

	return http.Post(getSendMessageURL(), "application/json", bytes.NewReader(body))
}
