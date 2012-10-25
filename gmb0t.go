package main

import (
	"time"
	"math/rand"
	"github.com/cmende/gmb0t/bot"
)

func main() {
	// we don't need a secure random
	rand.Seed(time.Now().UnixNano())

	bot.Connect()
	<-bot.Quit
}
