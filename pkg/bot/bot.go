package bot

import (
	"os"

	"github.com/shomali11/slacker"
)

var Bot *slacker.Slacker

func BotInit() {
	Bot = slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"), os.Getenv("SOCKET_TOKEN"))
}
