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

	switch cmd[0] {
	case "price":
		PriceHandler(s, m, cmd[1:])
	case "balance":
		BalanceHandler(s, m, cmd[1:])
	default:
		check(s.ChannelMessageSend(m.ChannelID, "What? Invalid command"))
	}
}

func commandParser(prefix, content string) ([]string, error) {
	var cmd []string
	if len(content) <= len(prefix) || content[len(prefix)] != ' ' {
		return cmd, fmt.Errorf("Comand %s invalid", content)
	}
	if content[:len(prefix)] == prefix {
		return strings.Fields(strings.TrimSpace(content[len(prefix):])), nil
	}
	return cmd, fmt.Errorf("Comand %s invalid", content)
}

func check(_ interface{}, err error) {
	if err != nil {
		fmt.Println("Error [message send]: ", err)
	}
}
