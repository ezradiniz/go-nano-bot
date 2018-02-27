package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/ezradiniz/go-nano-bot/handlers"
	"github.com/paked/configure"
)

var (
	conf   = configure.New()
	token  = conf.String("TOKEN", "token", "Bot token")
	prefix = conf.String("prefix", "?go", "Bot prefix")
)

func init() {
	conf.Use(configure.NewEnvironment())
	conf.Use(configure.NewFlag())
}

func main() {

	conf.Parse()

	discord, err := discordgo.New("Bot " + *token)
	if err != nil {
		fmt.Println("error creating Discord session", err)
		return
	}

	handler := handlers.NewBotHandler(*prefix)

	discord.AddHandler(handler.Commands)

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
