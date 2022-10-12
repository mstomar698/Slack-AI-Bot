package command

import (
	"github.com/mstomar698/Slack-AI-Bot/pkg/bot"
	"github.com/shomali11/slacker"
)

func Hello() {
	bot.BotInit()

	bot.Bot.Command("hello", &slacker.CommandDefinition{
		Description: "Say hello",
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			response.Reply("hello new user!")
		},
	})

}
