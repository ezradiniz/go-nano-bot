package util

import "github.com/bwmarrin/discordgo"

// MessageEmbedDefault create a message embed
func MessageEmbedDefault() *discordgo.MessageEmbed {

	// set thumbnail
	e := &discordgo.MessageEmbed{}
	e.Thumbnail = &discordgo.MessageEmbedThumbnail{URL: "https://i.imgur.com/Dc2rW20.png"}

	// set footer
	e.Footer = &discordgo.MessageEmbedFooter{Text: "Nano to the moon!"}

	e.URL = "https://github.com/ezradiniz/go-nano-bot"

	e.Color = 0x4a90e2

	return e
}

// AddFieldEmbed set field
func AddFieldEmbed(embed *discordgo.MessageEmbed, name, value string) {
	embed.Fields = append(embed.Fields, &discordgo.MessageEmbedField{
		Name:  name,
		Value: value,
	})
}
