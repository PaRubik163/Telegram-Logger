package config

import "os"

type Config struct {
	BotToken            string
	LogGroupID          string
	TopicUserID         string
	TopicExpenseID      string
	TopicIncomeID       string
	TopicSubscriptionID string
	GRPCPort string 
}

func NewConfig() *Config {
	return &Config{
		BotToken: os.Getenv("BOT_TOKEN"),
		LogGroupID: os.Getenv("LOG_GROUP_ID"),
		TopicUserID: os.Getenv("TOPIC_USER_ID"),
		TopicExpenseID: os.Getenv("TOPIC_EXPENSE_ID"),
		TopicIncomeID: os.Getenv("TOPIC_INCOME_ID"),
		TopicSubscriptionID: os.Getenv("TOPIC_SUBSCRIPTION_ID"),
		GRPCPort: os.Getenv("GRPC_PORT"),
	}
}