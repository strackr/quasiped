package main

import (
	"github.com/hybridgroup/gobot"
	"github.com/hybridgroup/gobot/platforms/raspi"

	"./quasiped"
)

func main() {
	gbot := gobot.NewGobot()
	r := raspi.NewRaspiAdaptor("raspi")
	gbot.AddRobot(quasiped.New(r, r).Robot())
	gbot.Start()
}
