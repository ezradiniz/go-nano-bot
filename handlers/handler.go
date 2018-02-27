package handlers

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
)

// BotHandler ...
type BotHandler struct {
	Prefix string
}

// NewBotHandler new NewBotHandler
func NewBotHandler(prefix string) *BotHandler {
	return &BotHandler{prefix}
}

// Commands handler method
func (h BotHandler) Commands(s *discordgo.Session, m *discordgo.MessageCreate) {
	if s.State.Ready.User.Username == m.Author.Username {
		return
	}

	cmd, err := commandParser(h.Prefix, m.Content)
	if err != nil {
		return
	}

	switch cmd {
	case "price":
		PriceHandler(s, m)
	default:
		check(s.ChannelMessageSend(m.ChannelID, "What? Invalid command"))
	}
}

func commandParser(prefix, content string) (string, error) {
	if len(content) <= len(prefix) || content[len(prefix)] != ' ' {
		return "", fmt.Errorf("Comand %s invalid", content)
	}
	if content[:len(prefix)] == prefix {
		return strings.TrimSpace(content[len(prefix):]), nil
	}
	return "", fmt.Errorf("Comand %s invalid", content)
}

func check(_ interface{}, err error) {
	if err != nil {
		fmt.Println("Error [message send]: ", err)
	}
}
