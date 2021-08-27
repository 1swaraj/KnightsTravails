package pkg

import "github.com/1swaraj/KnightTravails/utils"

type Chess struct {
	rows        int `yaml:"rows"`
	columns     int `yaml:"cols"`
	inputHelper utils.InputHelper
}

type Coordinates struct {
	x      int
	y      int
	parent *Coordinates
}