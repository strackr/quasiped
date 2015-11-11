package quasiped

import (
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
	joints map[location]*joint
}

func (q *Quasiped) Devices() []gobot.Device {
	var devs []gobot.Device
	for _, j := range q.joints {
		devs = append(devs, j.driver)
	}
	return devs
}

func New(w gpio.ServoWriter) *Quasiped {
	return &Quasiped{
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
