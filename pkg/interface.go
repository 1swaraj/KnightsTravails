package pkg

type IChess interface {
	CheckValidity(coordinates Coordinates) (error)
	PrintMoves(coordinates Coordinates)
}