package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Erro ao carregar o arquivo .env: %v", err)
	}
}

func GetSlackToken() string {
	token := os.Getenv("SLACK_BOT_TOKEN")
	if token == "" {
		log.Fatal("Variável SLACK_BOT_TOKEN não definida no .env")
	}
	return token
}

func GetSlackChannelID() string {
	channelID := os.Getenv("SLACK_CHANNEL_ID")
	if channelID == "" {
		log.Fatal("Variável SLACK_CHANNEL_ID não definida no .env")
	}
	return channelID
}
