package pkg

import "github.com/1swaraj/KnightTravails/utils"

type Chess struct {
	Rows        int `yaml:"board_rows"`
	Columns     int `yaml:"board_cols"`
	InputHelper utils.InputHelper `yaml:"input_helper"`
}

type Coordinates struct {
	x      int
	y      int
	parent *Coordinates
}