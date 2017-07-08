package main

import (
	"fmt"
	"os"
)

// Moves ...
const (
	Walk = 0
	Jump = 1
	High = 2
	Long = 3
)

func main() {
	if len(os.Args) != 2 {
		panic("Needs File")
	}

	// Board
	board, startX, startY := Build(os.Args[1])
	board.Display()
	fmt.Println(startX, startY)
}
