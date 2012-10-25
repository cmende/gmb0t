package bot

import (
	irc "github.com/fluffle/goirc/client"
)

var Quit chan bool = make(chan bool)
var commands chan []string = make(chan []string)

var server string = "irc.freenode.net"
var channel string = "##b0tgames"

func Connect() {
	c := irc.SimpleClient("gmb0t", "gmb0t", "Game Master")
	c.EnableStateTracking()
	c.AddHandler("connected", postConnect)
	c.AddHandler("disconnected", func(c *irc.Conn, l *irc.Line) { Quit <- true })
	c.AddHandler("NOTICE", parseNotice)
	go parseCommands()

	c.Connect(server)
}

func postConnect(conn *irc.Conn, line *irc.Line) {
	conn.Join(channel)
}

func parseNotice(conn *irc.Conn, line *irc.Line) {
	if line.Args[0] == conn.Me.Nick {
		commands <- []string{line.Nick, line.Args[1]}
	}
}
