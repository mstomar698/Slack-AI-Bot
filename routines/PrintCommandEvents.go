package routines

import (
	"fmt"

	"github.com/shomali11/slacker"
)

func PrintCommandEvents(analyticsChannel <-chan *slacker.CommandEvent) {
	for event := range analyticsChannel {
		fmt.Println("Command Events")
		fmt.Println(event.Timestamp)
		fmt.Println(event.Command)
		// fmt.Println(event.User)
		// fmt.Println(event.Channel)
		// fmt.Println(event.Team)
		fmt.Println(event.Parameters)
		fmt.Println(event.Event)
		fmt.Println()
	}
}

