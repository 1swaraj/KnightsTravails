package pkg

import (
	"bufio"
	"fmt"
	"github.com/1swaraj/KnightTravails/utils"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

const DefaultBoardSize = 8

// Messages
const EndingPosition = "Enter the Ending Position :"
const StartingPoisition = "Enter the Starting Position :"
const OutofBoundary = "The position is out of boundary"
const CouldNotReadYaml = "Could not find your config.yaml file so using the default board size"
const KnightCantMove = "The knight cannot move from the starting position to the ending position"

// Valid moves for a knight
var moves = []Coordinates{{2, 1, nil}, {2, -1, nil}, {1, -2, nil}, {-1, -2, nil}, {-2, 1, nil}, {-2, -1, nil}, {1, 2, nil}, {-1, 2, nil}}

// global variables
var chess Chess
var startingCoordinates Coordinates
var endingCoordinates Coordinates

// Initializing the chessboard by reading the config from conf.yaml
func InitConfigs() {
	yamlFile, err := ioutil.ReadFile("./config.yaml")
	if err != nil {
		fmt.Println(CouldNotReadYaml)
		chess = Chess{Rows: DefaultBoardSize, Columns: DefaultBoardSize, InputHelper: utils.InputHelper{bufio.NewScanner(os.Stdin)}}
		return
	}
	err = yaml.Unmarshal(yamlFile, &chess)
	if err != nil {
		fmt.Println(CouldNotReadYaml)
		chess = Chess{Rows: DefaultBoardSize, Columns: DefaultBoardSize, InputHelper: utils.InputHelper{bufio.NewScanner(os.Stdin)}}
		return
	}
	chess.InputHelper = utils.InputHelper{bufio.NewScanner(os.Stdin)}
}
