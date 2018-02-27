package handlers

import (
	"fmt"
	"strconv"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/ezradiniz/go-nano-bot/api"
	"github.com/ezradiniz/go-nano-bot/util"
)

// PriceHandler ...
func PriceHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	nanoResponse := api.NanoResponse{}

	err := api.FetchNano(&nanoResponse)
	if err != nil {
		fmt.Println("Error:", err)
		check(s.ChannelMessageSend(m.ChannelID, "Sorry, an error ocurred while processing your request"))
	}

	nano := nanoResponse[0]

	embed := util.MessageEmbedDefault()

	// set Author
	embed.Author = &discordgo.MessageEmbedAuthor{
		Name: "CoinMarketCap",
		URL:  "https://coinmarketcap.com/currencies/nano/",
	}

	util.AddFieldEmbed(embed, "Price USD", nano.PriceUSD)
	util.AddFieldEmbed(embed, "Price BTC", nano.PriceBTC)
	util.AddFieldEmbed(embed, "Rank", nano.Rank)

	// parse time
	timeInt, _ := strconv.ParseInt(nano.LastUpdated, 10, 64)
	timeUnix := time.Unix(timeInt, 0)
	util.AddFieldEmbed(embed, "Last Updated", timeUnix.String())

	check(s.ChannelMessageSendEmbed(m.ChannelID, embed))
}
