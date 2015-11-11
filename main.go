package main

import (
	"log"
	"time"

	"github.com/hybridgroup/gobot"
	"github.com/hybridgroup/gobot/platforms/raspi"

	"./quasiped"
)

func main() {
	gbot := gobot.NewGobot()

	r := raspi.NewRaspiAdaptor("raspi")
	q := quasiped.New(r)

	work := func() {
		gobot.Every(1*time.Second, func() {
			log.Printf("Update loop %s", time.Now().String())
		})
	}

	robot := gobot.NewRobot("quasiped",
		[]gobot.Connection{r},
		q.Devices(),
		work,
	)
	gbot.AddRobot(robot)

	gbot.Start()
}
