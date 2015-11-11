package quasiped

import (
	"log"
	"time"

	"github.com/hybridgroup/gobot"
	"github.com/hybridgroup/gobot/platforms/gpio"
)

type position int

const (
	nw position = iota
	ne
	sw
	se
)

type orientation int

const (
	horizontal orientation = iota
	vertical
)

type location struct {
	position    position
	orientation orientation
}

type joint struct {
	driver *gpio.ServoDriver
}

type Quasiped struct {
	connection gobot.Connection
	joints     map[location]*joint
}

func (q *Quasiped) Robot() *gobot.Robot {
	work := func() {
		gobot.Every(1*time.Second, func() {
			log.Printf("Update loop %s", time.Now().String())
		})
	}

	return gobot.NewRobot("quasiped",
		[]gobot.Connection{q.connection},
		q.devices(),
		work,
	)
}

func (q *Quasiped) devices() []gobot.Device {
	var devs []gobot.Device
	for _, j := range q.joints {
		devs = append(devs, j.driver)
	}
	return devs
}

func New(c gobot.Connection, w gpio.ServoWriter) *Quasiped {
	return &Quasiped{
		connection: c,
		joints: map[location]*joint{
			location{nw, horizontal}: &joint{driver: gpio.NewServoDriver(w, "nw_h", "1")},
			location{ne, horizontal}: &joint{driver: gpio.NewServoDriver(w, "ne_h", "2")},
			location{sw, horizontal}: &joint{driver: gpio.NewServoDriver(w, "sw_h", "3")},
			location{se, horizontal}: &joint{driver: gpio.NewServoDriver(w, "se_h", "4")},
			location{nw, vertical}:   &joint{driver: gpio.NewServoDriver(w, "nw_v", "5")},
			location{ne, vertical}:   &joint{driver: gpio.NewServoDriver(w, "ne_v", "6")},
			location{sw, vertical}:   &joint{driver: gpio.NewServoDriver(w, "sw_v", "7")},
			location{se, vertical}:   &joint{driver: gpio.NewServoDriver(w, "se_v", "8")},
		},
	}
}
