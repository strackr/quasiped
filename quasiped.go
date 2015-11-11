package main

import (
	"log"
	"time"

	"github.com/hybridgroup/gobot"
	"github.com/hybridgroup/gobot/platforms/raspi"
)

func main() {
	gbot := gobot.NewGobot()

	r := raspi.NewRaspiAdaptor("raspi")

	work := func() {
		gobot.Every(1*time.Second, func() {
			log.Printf("Update loop %s", time.Now().String())
		})
	}

	robot := gobot.NewRobot("quasiped",
		[]gobot.Connection{r},
		[]gobot.Device{},
		work,
	)
	gbot.AddRobot(robot)

	gbot.Start()
}
