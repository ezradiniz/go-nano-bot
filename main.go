package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/ezradiniz/go-nano/handlers"
)

func main() {

	discord, err := discordgo.New("Bot " + os.Getenv("TOKEN"))
	if err != nil {
		fmt.Println("error creating Discord session", err)
		return
	}

	discord.AddHandler(handlers.BotHandler)

	err = discord.Open()
	if err != nil {
		fmt.Println("error opening connection", err)
		return
	}

	fmt.Println("Bot is running. Press CTRL-C to exit")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)

	<-sc

	if err := discord.Close(); err != nil {
		fmt.Println("Failed to exit")
	}
}
