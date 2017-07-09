package main

import (
	"fmt"
	"math/rand"
	"os"
)

// Random returns a new random board and the starting coordinates
func Random(width, height int, density float64, seed int) *Board {

	// Seed
	rand.Seed(int64(seed))

	// New Board
	b := NewBoard(width, height)

	// Fill Randomly
	for x := 1; x < width; x++ {
		for y := 0; y < height; y++ {
			if rand.Intn(1000) < int(density*1000) {
				if rand.Int()%2 == 0 {
					b.Add(x, y, Wall)
				} else {
					b.Add(x, y, Coin)
				}
			}
		}
	}

	return b
}

// Build returns a new board and the starting coordinates
func Build(filename string) *Board {

	// Open File
	file, err := os.Open(filename)
	dieOn(err)

	// Read Dimensions
	var width int
	var height int
	fmt.Fscanf(file, "%d %d\n", &width, &height)

	// Read Starting
	var startX int
	var startY int
	fmt.Fscanf(file, "%d %d\n", &startX, &startY)

	// New Board
	b := NewBoard(width, height)
	b.StartX = startX
	b.StartY = startY

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

	return b
}
