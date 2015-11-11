package quasiped

import (
	"github.com/hybridgroup/gobot/platforms/gpio"
)

type position int

const (
	nw position = iota
	ne position = iota
	sw position = iota
	se position = iota
)

type orientation int

const (
	horizontal orientation = iota
	vertical   orientation = iota
)

type location struct {
	position    position
	orientation orientation
}

type engine struct {
	servos map[location]*servo
}

type servo struct {
	driver *gpio.ServoDriver
}

func New(w gpio.ServoWriter) *engine {
	return &engine{
		servos: map[location]*servo{
			location{nw, horizontal}: &servo{driver: gpio.NewServoDriver(w, "nw_h", "1")},
			location{ne, horizontal}: &servo{driver: gpio.NewServoDriver(w, "ne_h", "2")},
			location{sw, horizontal}: &servo{driver: gpio.NewServoDriver(w, "sw_h", "3")},
			location{se, horizontal}: &servo{driver: gpio.NewServoDriver(w, "se_h", "4")},
			location{nw, vertical}:   &servo{driver: gpio.NewServoDriver(w, "nw_v", "5")},
			location{ne, vertical}:   &servo{driver: gpio.NewServoDriver(w, "ne_v", "6")},
			location{sw, vertical}:   &servo{driver: gpio.NewServoDriver(w, "sw_v", "7")},
			location{se, vertical}:   &servo{driver: gpio.NewServoDriver(w, "se_v", "8")},
		},
	}
}
