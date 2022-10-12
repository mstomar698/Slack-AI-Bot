package main

import (
	"context"
	"log"
	"os"

	"github.com/mstomar698/Slack-AI-Bot/pkg/bot"
	"github.com/mstomar698/Slack-AI-Bot/pkg/command"
	"github.com/mstomar698/Slack-AI-Bot/routines"
	"github.com/spf13/viper"
)

func viperEnvVariable(key string) string {
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error while reading config file %s", err)
	}
	value, ok := viper.Get(key).(string)
	if !ok {
		log.Fatalf("Invalid type assertion")
	}
	return value
}

func main() {
	slackbotkey := viperEnvVariable("SLACK_BOT_TOKEN")
	sockettokenkey := viperEnvVariable("SOCKET_TOKEN")
	// Setting up env variables
	os.Setenv("SLACK_BOT_TOKEN", slackbotkey)
	os.Setenv("SOCKET_TOKEN", sockettokenkey)
	bot.BotInit()
	// Calling commands only one command could be run at a time
	command.Hello()

	go routines.PrintCommandEvents(bot.Bot.CommandEvents())

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	err := bot.Bot.Listen(ctx)
	if err != nil {
		log.Fatal(err)
	}
}
