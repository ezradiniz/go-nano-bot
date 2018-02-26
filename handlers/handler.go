package handlers

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
)

var prefix = "?go"

func command(content string) string {
	if len(content) <= len(prefix) {
		return ""
	}
	if content[:len(prefix)] == prefix {
		return strings.TrimSpace(content[len(prefix):])
	}
	return ""
}

func errHandler(_ interface{}, err error) {
	if err != nil {
		fmt.Println("Error [message send]: ", err)
	}
}

// BotHandler handler method
func BotHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if s.State.Ready.User.Username == m.Author.Username {
		return
	}
	cmd := command(m.Content)
	switch cmd {
	}
}

// SetPrefix set new prefix
func SetPrefix(p string) {
	prefix = p
}
