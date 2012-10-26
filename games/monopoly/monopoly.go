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
var currentPlayer int = 0

func StartGame(nicks []string) {
	for _, nick := range nicks {
		fmt.Printf("New player: %s\n", nick)
		players = append(players, Player{nick: nick, money: 1500, position: 0})
	}
	nextPlayer()
}

func nextPlayer() {
	player := &players[currentPlayer]
	result := roll()
	player.position += result
	bot.Say(fmt.Sprintf("%s, it's your turn. You rolled %d. New position: %d", player.nick, result, player.position))

	currentPlayer++
	if currentPlayer >= len(players) {
		currentPlayer = 0
	}
}

func Parse(sender, action string) {
	switch action {
	case "done":
		nextPlayer()
	default:
	}
}

func roll() int {
	return rand.Intn(6) + 1 + rand.Intn(6) + 1
}
