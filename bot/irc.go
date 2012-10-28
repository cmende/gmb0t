package bot

import (
	irc "github.com/fluffle/goirc/client"
)

var (
	Quit = make(chan bool)
	commands = make(chan []string)

	server = "irc.freenode.net"
	channel = "##b0tgames"
	conn *irc.Conn
)

func Connect() {
	conn = irc.SimpleClient("gmb0t", "gmb0t", "Game Master")
	conn.EnableStateTracking()
	conn.AddHandler("connected", postConnect)
	conn.AddHandler("disconnected", func(c *irc.Conn, l *irc.Line) { Quit <- true })
	conn.AddHandler("NOTICE", parseNotice)
	go parseCommands()

	conn.Connect(server)
}

func postConnect(conn *irc.Conn, line *irc.Line) {
	conn.Join(channel)
}

func parseNotice(conn *irc.Conn, line *irc.Line) {
	if line.Args[0] == conn.Me.Nick {
		commands <- []string{line.Nick, line.Args[1]}
	}
}

func Say(text string) {
	conn.Privmsg(channel, text)
}

func Notice(target, text string) {
	conn.Notice(target, text)
}
