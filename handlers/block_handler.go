package handlers

import (
	"github.com/bwmarrin/discordgo"
	"github.com/ezradiniz/go-nano-bot/api"
	"github.com/ezradiniz/go-nano-bot/util"
)

// BlockHandler ...
func BlockHandler(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {
	block := api.Block{}

	err := api.FetchBlock(&block)
	if err != nil {
		check(s.ChannelMessageSend(m.ChannelID, "Sorry, an error ocurred while processing your request"))
	}

	embed := util.MessageEmbedDefault()
	embed.Author = &discordgo.MessageEmbedAuthor{
		Name: "Nanode.co",
		URL:  "https://www.nanode.co",
	}

	util.AddFieldEmbed(embed, "Blocks", block.Value)

	check(s.ChannelMessageSendEmbed(m.ChannelID, embed))
}
