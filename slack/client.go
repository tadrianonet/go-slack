package slack

import (
	"fmt"
	"log"

	"github.com/slack-go/slack"
)

type Client struct {
	api       *slack.Client
	channelID string
}

func NewClient(token, channelID string) *Client {
	return &Client{
		api:       slack.New(token),
		channelID: channelID,
	}
}

func (c *Client) SendMessage(message string) {
	_, _, err := c.api.PostMessage(
		c.channelID,
		slack.MsgOptionText(message, false),
	)
	if err != nil {
		log.Fatalf("Erro ao enviar mensagem: %v", err)
	}
	fmt.Println("Mensagem enviada com sucesso!")
}
