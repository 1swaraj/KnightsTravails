package utils

import (
	"fmt"
	"strconv"
	"strings"
)


func (inputHelper *InputHelper) GetPosition(msg string) (row int,column int,err error) {
	fmt.Println(msg)
	var algebraicNotation string
	_,err = fmt.Scanf("%s",&algebraicNotation)
	if err != nil {
		return
	}
	row,column,err = getCoordinates(algebraicNotation)
	// Retry incase of invalid input
	if err !=nil {
		fmt.Println(err.Error())
		return inputHelper.GetPosition(msg)
	}
	return row,column,err
}

func getCoordinates(algebraicNotion string) (row int, column int, err error) {
	algebraicNotion = strings.ToUpper(algebraicNotion)
	// ASCII approach for getting the rows
	// A = 65
	row = int(algebraicNotion[0] - 65)
	column, err = strconv.Atoi(string(algebraicNotion[1:]))
	if err != nil {
		err = fmt.Errorf("Invalid Position")
		return
	}
	// Need to subtract so that we maintain parity with index of our
	// arrays starting from 0
	column --
	// There are only 26 letters so the maximum number of columns allowed will be 26
	// This also helps us in detecting that the first character is a letter
	// because we have taken the ascii approach
	// This is just to check if the entered value is a valid algebraic notation
	// The actual validation (position out of board) is done in the CheckValidity method
	if row < 0 && row >= 26 {
		err = fmt.Errorf("Invalid Position")
		return
	}
	return
}

