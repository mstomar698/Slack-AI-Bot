package command

import (
	"github.com/mstomar698/Slack-AI-Bot/pkg/bot"
	"github.com/shomali11/slacker"
)

func Hello() {
	bot.BotInit()

	bot.Bot.Command("hello - <user>", &slacker.CommandDefinition{
		Description: "Say hello to someone",
		// Example:	 "hello - user",
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			query := request.Param("user")
			res := "Welcome to dev-lok" + "\n 👋 Hello From  " + query + "  and others"
			response.Reply(res)
		},
	})

}
