package main

import (
	"github.com/cmende/gmb0t/bot"
	"github.com/cmende/gmb0t/games/monopoly"
	"math/rand"
	"time"
)

func main() {
	// we don't need a secure random
	rand.Seed(time.Now().UnixNano())

	// load games
	bot.Games = append(bot.Games, bot.GameInfo{Name: "Monopoly", StartCallback: monopoly.StartGame, ActionCallback: monopoly.Parse})

	bot.Connect()
	<-bot.Quit
}
