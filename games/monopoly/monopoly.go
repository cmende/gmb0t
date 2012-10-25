package monopoly

import (
	"github.com/cmende/gmb0t/bot"
	"fmt"
	"math/rand"
)

func StartGame(players []string) {
	roll1, roll2 := roll()
	bot.Say(fmt.Sprintf("%d/%d", roll1, roll2))
}

func Parse(sender, action string) {
}

func roll() (int, int) {
	return rand.Intn(6) + 1, rand.Intn(6) + 1
}
