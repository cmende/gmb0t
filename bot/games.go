package bot

import (
	"fmt"
	"math/rand"
)

type GameInfo struct {
	name           string
	startCallback  func(players []string)
	actionCallback func(sender, action string)
}

var games []GameInfo
var currentGame GameInfo
var gameActive bool = false
var players []string

func RegisterGame(game GameInfo) {
	games = append(games, game)
}

func startGame() {
	if len(players) < 2 {
		fmt.Println("Error: not enough players")
		return
	}
	if len(games) < 1 {
		fmt.Println("Error: no games loaded")
		return
	}

	game := rand.Intn(len(games))
	currentGame = games[game]
	fmt.Printf("Starting game: %s\n", currentGame.name)
	currentGame.startCallback(players)
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
		case "register":
			players = append(players, nick)
			if gameActive {
				Notice(nick, "ok gameinprogress")
			} else {
				Notice(nick, "ok")
			}
		default:
			if gameActive {
				go currentGame.actionCallback(nick, command)
			} else {
				fmt.Println("Error: no current game")
			}
		}
	}
}
