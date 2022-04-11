# twitchbot
Go Twitch Bot Api wrapper, with an easy to use interface.

# Example
```go
package main

import (
	"github.com/witer33/twitchbot"
)

func main() {
	bot := twitchbot.NewBot("oauth:abcdef", "mybot", []string{"channel"})

	bot.OnMessage(func(bot *twitchbot.Bot, message *twitchbot.Message) {
		if message.Message == "!ping" {
			message.Reply("pong")
			message.Delete()
		}
	})

	bot.Run()
}
```
