package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"twitch/twitchbot"
)

type Config struct {
	Bot struct {
		Token string `json:"token"`
		Nick  string `json:"nick"`
	} `json:"bot"`
}

func main() {

	var config Config

	data, _ := ioutil.ReadFile("config.json")
	err := json.Unmarshal(data, &config)
	if err != nil {
		log.Panicln(err)
	}

	bot := twitchbot.NewBot(config.Bot.Token, config.Bot.Nick, []string{"witer33"})

	bot.OnMessage(func(bot *twitchbot.Bot, message *twitchbot.Message) {
		fmt.Println(message)
		if message.Message == "!ping" {
			message.Reply("pong")
			message.Delete()
		}
	})

	bot.Run()
}
