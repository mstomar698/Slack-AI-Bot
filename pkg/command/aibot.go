package command

import (
	"encoding/json"
	"fmt"

	"github.com/krognol/go-wolfram"
	"github.com/mstomar698/Slack-AI-Bot/pkg/bot"
	"github.com/mstomar698/Slack-AI-Bot/pkg/client"
	"github.com/shomali11/slacker"
	"github.com/tidwall/gjson"
	witai "github.com/wit-ai/wit-go/v2"
)

func Aibot() {
	// calling bot & client from pkg
	bot.BotInit()
	client.ClientInit()
	// This is where the AI bot will be called
	bot.Bot.Command("Query for AI-Bot - <msg>", &slacker.CommandDefinition{
		Description: "Send any question to wolfram",
		// Example:	 "What is the weather in New York?",
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			// this is how you pass param in go
			query := request.Param("msg")
			message, _ := client.Client.Parse(&witai.MessageRequest{
				Query: query,
			})
			data, _ := json.MarshalIndent(message, "", "    ")
			rough := string(data[:])
			// gjson is a json parser which allows us to use json as object just like in js
			value := gjson.Get(rough, "entities.wit$wolfram_search_query:wolfram_search_query.0.value" )
			// we require value from json to be passed to wolfram
			answer := value.String()
			// pass answer to wolfram
			res, err := client.WolframClient.GetSpokentAnswerQuery(answer, wolfram.Metric, 1000)
			if err != nil {
				fmt.Println("error have occured")
			}
			fmt.Println(value)
			response.Reply(res)
		},
	})
}