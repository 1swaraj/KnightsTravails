package utils

type IInputHelper interface {
	GetPosition(string) (int, int, error)
}
