package handlers

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/ezradiniz/go-nano-bot/api"
	"github.com/ezradiniz/go-nano-bot/util"
)

const qrCode = "https://chart.googleapis.com/chart?chs=200x200&cht=qr&chl="

// BalanceHandler ...
func BalanceHandler(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {
	if len(args) == 0 {
		check(s.ChannelMessageSend(m.ChannelID, "Please inform your address"))
		return
	}

	balance := api.Balance{}

	err := api.FetchBalance(args[0], &balance)

	if err != nil {
		fmt.Println("Error:", err)
		check(s.ChannelMessageSend(m.ChannelID, "Sorry, an error ocurred while processing your request"))
		return
	}

	if balance.Amount == "" {
		check(s.ChannelMessageSend(m.ChannelID, "Invalid address"))
		return
	}

	embed := util.MessageEmbedDefault()

	embed.Author = &discordgo.MessageEmbedAuthor{
		Name: "Address",
		URL:  api.Raiblocks + "?acc=" + args[0],
	}

	embed.Image = &discordgo.MessageEmbedImage{
		URL: qrCode + args[0],
	}

	util.AddFieldEmbed(embed, "Amount", balance.Amount)

	check(s.ChannelMessageSendEmbed(m.ChannelID, embed))
}
