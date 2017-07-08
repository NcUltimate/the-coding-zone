package main

import (
	"os"
)

func main() {
	if len(os.Args) != 2 {
		panic("Needs File")
	}

	// Board
	board, startX, startY := Build(os.Args[1])
	board.Display()
	Solve(board, startX, startY)
}
