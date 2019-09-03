package main

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

var discord *discordgo.Session

func main() {
	var err error
	if err = loadConf(); err != nil {
		log.Fatalln(err)
	}

	discord, err = discordgo.New("Bot " + conf.BotToken)
	if err != nil {
		log.Fatalln(err)
	}

	discord.AddHandler(inputHandler)

	if err = discord.Open(); err != nil {
		log.Fatalln(err)
	}
	defer discord.Close()

	<-make(chan struct{})
}
