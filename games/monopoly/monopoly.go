package monopoly

import (
	"fmt"
	"github.com/cmende/gmb0t/bot"
	"math/rand"
)

type Player struct {
	nick            string
	money, position int
}

var players []Player

func StartGame(nicks []string) {
	for _, nick := range nicks {
		fmt.Printf("New player: %s\n", nick)
		players = append(players, Player{nick: nick, money: 1500, position: 0})
	}
	roll1, roll2 := roll()
	bot.Say(fmt.Sprintf("%d/%d", roll1, roll2))
}

func Parse(sender, action string) {
}

func roll() (int, int) {
	return rand.Intn(6) + 1, rand.Intn(6) + 1
}
