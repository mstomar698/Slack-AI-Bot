package main

import (
	"context"
	"log"

	"github.com/joho/godotenv"
	"github.com/mstomar698/Slack-AI-Bot/pkg/bot"
	"github.com/mstomar698/Slack-AI-Bot/pkg/client"
	"github.com/mstomar698/Slack-AI-Bot/pkg/command"
	"github.com/mstomar698/Slack-AI-Bot/routines"
)

func main() {
	// Loading variable using godotenv
	godotenv.Load(".env")

	// calling bot & client from pkg
	bot.BotInit()
	client.ClientInit()
	
	// Calling commands only one command could be run at a time
	// command.Aibot()
	command.Hello()

	go routines.PrintCommandEvents(bot.Bot.CommandEvents())

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	err := bot.Bot.Listen(ctx)
	if err != nil {
		log.Fatal(err)
	}
}
