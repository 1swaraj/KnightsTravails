package pkg

import (
	"bufio"
	"github.com/1swaraj/KnightTravails/utils"
	"github.com/spf13/viper"
	"os"
)

const DefaultBoardSize = 8

// Messages
const EndingPosition = "Enter the Ending Position :"
const StartingPoisition = "Enter the Starting Position :"
const OutofBoundary = "The position is out of boundary"
const CouldNotReadYaml = "Could not find your config file so using the default board size"
const KnightCantMove = "The knight cannot move from the starting position to the ending position"

// Valid moves for a knight
var moves = []Coordinates{{2, 1, nil}, {2, -1, nil}, {1, -2, nil}, {-1, -2, nil}, {-2, 1, nil}, {-2, -1, nil}, {1, 2, nil}, {-1, 2, nil}}

// global variables
var chess Chess
var startingCoordinates Coordinates
var endingCoordinates Coordinates

// Initializing the chessboard by reading the config from conf.yaml
func InitConfigs() {
	rows := viper.GetInt("board_rows")
	cols := viper.GetInt("board_cols")
	if rows == 0 || cols == 0 {
		rows = DefaultBoardSize
		cols = DefaultBoardSize
	}
	chess = Chess{Rows: rows, Columns: cols, InputHelper: utils.InputHelper{bufio.NewScanner(os.Stdin)}}
	return
}
