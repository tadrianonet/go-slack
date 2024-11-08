package main

import (
	"go-slack/config"
	"go-slack/slack"
)

func main() {
	config.LoadEnv()

	token := config.GetSlackToken()
	channelID := config.GetSlackChannelID()

	client := slack.NewClient(token, channelID)

	client.SendMessage("Olá, Slack! Esta é uma mensagem enviada pelo Go.")
}
