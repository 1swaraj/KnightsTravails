# KnightsTravails

### Understanding the Code Structure

Maintaining the standard directory structure for golang applications

##### pkg

###### config.go
This file contains the strings (messages, error messages etc), the variables needed for our operations (chessBoard, startingPosition, endingPosition)
Also defines the valid moves for the knight. (Extensibility aspect)
It has the InitConfig method which reads the configs from a yaml file. Making this solution extendable

###### interface.go
Contains the IChess Interface
This interface is implementated to act as a blueprint but more importantly to make testing easy as the project size increases.

###### knight.go
Contains the core logic of the Knights Travails Problem

###### types.go
Contians the structs needed for implementing the solution.
struct Chess maintains the state of the chess board - rows and columns
struct Chess also has a field for input helper which will help us in testing (mocks)
struct Coordinates has  

##### utils

###### interface.go
IInputHelper is an interface that serves as a blueprint for IO related functionalities

###### types.go
Defines the InputHelper struct which has a field for IO
This is done so that if in the future we decide to change the implementation of IO from buffio to something else, 
we won't have to change the logic in the calling functions

###### utils.go
Defines all utility functionalities such as getting user input

### Instructions for usage
```
git clone https://github.com/1swaraj/KnightsTravails
cd KnightsTravails
go mod tidy
go run .
```

If you want to change the size of the board just edit config.yaml