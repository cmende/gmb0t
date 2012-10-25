package bot

import (
	"math/rand"
	"fmt"
)

type GameInfo struct {
	name string
	startCallback func()
	actionCallback func(sender, action string)
}

var games []GameInfo
var currentGame GameInfo
var gameActive bool = false

func RegisterGame(game GameInfo) {
	games = append(games, game)
}

func startGame() {
	if (len(games) < 1) {
		fmt.Println("Error: no games loaded")
		return
	}

	game := rand.Intn(len(games))
	currentGame = games[game]
	fmt.Printf("Starting game: %s\n", currentGame.name)
	currentGame.startCallback()
	gameActive = true
}

func parseCommands() {
	for {
		line := <-commands
		nick := line[0]
		command := line[1]
		fmt.Printf("<%s> %s\n", nick, command)

		switch command {
		case "start":
			startGame()
		default:
			if gameActive {
				go currentGame.actionCallback(nick, command)
			} else {
				fmt.Println("Error: no current game")
			}
		}
	}
}
