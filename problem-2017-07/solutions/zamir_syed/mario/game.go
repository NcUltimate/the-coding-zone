package main

import (
	"fmt"
	"os"
)

// Build returns a new board and the starting coordinates
func Build(filename string) (*Board, int, int) {

	// Open File
	file, err := os.Open(filename)
	dieOn(err)

	// Read Dimensions
	var width int
	var height int
	fmt.Fscanf(file, "%d %d\n", &width, &height)

	var startX int
	var startY int
	fmt.Fscanf(file, "%d %d\n", &startX, &startY)

	// New Board
	b := NewBoard(width, height)

	// Add Objects
	for {
		var x int
		var y int
		var o byte

		_, err := fmt.Fscanf(file, "%d %d %c\n", &x, &y, &o)
		if err != nil {
			break
		}

		switch {
		case o == 'C':
			b.Add(x, y, Coin)
		case o == 'B':
			b.Add(x, y, Wall)
		}
	}

	return b, startX, startY
}
