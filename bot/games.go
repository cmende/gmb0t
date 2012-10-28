package bot

import (
	"fmt"
	"math/rand"
)

type GameInfo struct {
	Name           string
	StartCallback  func(players []string)
	ActionCallback func(sender, action string)
}

var (
	Games []GameInfo
	currentGame GameInfo
	gameActive = false
	players []string
)

func RegisterGame(game GameInfo) {
	Games = append(Games, game)
}

func startGame() {
	if len(players) < 2 {
		fmt.Println("Error: not enough players")
		return
	}
	if len(Games) < 1 {
		fmt.Println("Error: no games loaded")
		return
	}

	game := rand.Intn(len(Games))
	currentGame = Games[game]
	fmt.Printf("Starting game: %s\n", currentGame.Name)

	// shuffle players before starting (Fisher-Yates shuffle)
	for i := range players {
		j := rand.Intn(i + 1)
		players[i], players[j] = players[j], players[i]
	}

	currentGame.StartCallback(players)
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
				go currentGame.ActionCallback(nick, command)
			} else {
				fmt.Println("Error: no current game")
			}
		}
	}
}
