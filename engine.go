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

type Engine struct {
	servos map[location]*gpio.ServoDriver
}

func NewEngine(w gpio.ServoWriter) *Engine {
	return &Engine{
		servos: map[location]*gpio.ServoDriver{
			location{nw, horizontal}: gpio.NewServoDriver(w, "nw_h", "1"),
			location{ne, horizontal}: gpio.NewServoDriver(w, "ne_h", "2"),
			location{sw, horizontal}: gpio.NewServoDriver(w, "sw_h", "3"),
			location{se, horizontal}: gpio.NewServoDriver(w, "se_h", "4"),
			location{nw, vertical}:   gpio.NewServoDriver(w, "nw_v", "5"),
			location{ne, vertical}:   gpio.NewServoDriver(w, "ne_v", "6"),
			location{sw, vertical}:   gpio.NewServoDriver(w, "sw_v", "7"),
			location{se, vertical}:   gpio.NewServoDriver(w, "se_v", "8"),
		},
	}
}
