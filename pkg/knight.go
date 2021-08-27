package pkg

import (
	"fmt"
)

func Init() {
	InitConfigs()
	startingCoordinates = chess.GetInputs(StartingPoisition)
	endingCoordinates = chess.GetInputs(EndingPosition)
	chess.FindPath(startingCoordinates, endingCoordinates, chess.rows, chess.columns)
}

// Checks if the given coordinates are valid or not
func (chess *Chess) CheckValidity(coordinates Coordinates) (err error) {
	if coordinates.x > chess.columns-1 || coordinates.x < 0 || coordinates.y > chess.rows-1 || coordinates.y < 0 {
		return fmt.Errorf(OutofBoundary)
	}
	return
}

// Getting the input and checking the validity of the input
func (chess *Chess) GetInputs(msg string) Coordinates {
	x, y, err := chess.inputHelper.GetPosition(msg)
	if err != nil {
		fmt.Println(err.Error())
	}
	coordinates := Coordinates{x: x, y: y}
	err = chess.CheckValidity(coordinates)
	if err != nil {
		fmt.Println(err.Error())
		return chess.GetInputs(msg)
	}
	return coordinates
}

// Simple BFS implementation for solving the Knights Travails Problem
func (chess *Chess) FindPath(source Coordinates, dest Coordinates, r, c int) {
	visited := make(map[Coordinates]bool, r*c)
	var queue []Coordinates
	queue = append(queue, source)
	n := len(queue)
	for n != 0 {
		position := queue[n-1]
		queue = queue[:n-1]
		if position.x == dest.x && position.y == dest.y {
			chess.PrintMoves(position)
			return
		}
		// Skip if the node is already visited
		if _, ok := visited[position]; !ok {
			// ignoring the pointer reference
			visited[position] = true
			for _, move := range moves {
				potentialMove := Coordinates{x: position.x + move.x, y: position.y + move.y}
				if err := chess.CheckValidity(potentialMove); err == nil {
					potentialMove.parent = &position
					queue = append([]Coordinates{potentialMove}, queue...)
				}
			}
		}
		n = len(queue)
	}
}

// Printing the shortest path for the knight
func (chess *Chess) PrintMoves(coordinates Coordinates) {
	moves := ""
	movesNormal := ""
	for coordinates.parent != nil {
		movesNormal = fmt.Sprintf("{%d,%d} ", coordinates.x, coordinates.y) + movesNormal
		moves = fmt.Sprintf("%s%d ", string(coordinates.x+65), coordinates.y+1) + moves
		coordinates = *coordinates.parent
	}
	fmt.Println(moves)
}
