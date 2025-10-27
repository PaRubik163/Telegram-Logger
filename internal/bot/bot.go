package bot

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"teleglogger/internal/config"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Bot struct{
	bot *tgbotapi.BotAPI
	config *config.Config
}

func NewBot(conf *config.Config) (*Bot, error) {
	bot, err := tgbotapi.NewBotAPI(conf.BotToken)

	if err != nil{
		return nil, fmt.Errorf("failed to create new bot %v", err)
	}

	return &Bot{
		bot: bot,
		config: conf,
	}, nil
}

func (b *Bot) Send(topic, level, text string) error {
	topicID := map[string]string{
		"Users": b.config.TopicUserID,
		"Expense": b.config.TopicExpenseID,
		"Income": b.config.TopicIncomeID,
		"Subscription": b.config.TopicSubscriptionID,
	}[topic]

	emoji := map[string]string{
		"INFO": "1",
		"WARN": "2",
		"ERROR": "3",
	}[level]

	msg := fmt.Sprintf("%s [%s] %s", emoji, level, text)

	payload := map[string]any{
		"chat_id": b.config.LogGroupID,
		"message_thread_id": topicID,
		"text": msg,
	}

	data, err := json.Marshal(payload)

	if err != nil{
		return fmt.Errorf("failed to marshal payload JSON %v", err)
	}

	url := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", b.config.BotToken)

	_, err = http.Post(url, "application/json", bytes.NewReader(data))

	if err != nil{
		return fmt.Errorf("failed to send message POST %v", err)
	}

	return nil
}