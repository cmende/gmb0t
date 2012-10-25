package main

import (
	"github.com/cmende/gmb0t/bot"
	"math/rand"
	"time"
)

func main() {
	// we don't need a secure random
	rand.Seed(time.Now().UnixNano())

	bot.Connect()
	<-bot.Quit
}
