package main

import (
	"fmt"
	"os"
)

// Objects ...
const (
	None = 0
	Coin = 1
	Wall = 2
)

// Board ...
type Board struct {
	W      int
	H      int
	StartX int
	StartY int
	Spaces [][]int
}

// NewBoard ...
func NewBoard(w, h int) *Board {
	b := &Board{W: w, H: h}
	b.Spaces = make([][]int, w)
	for x := 0; x < w; x++ {
		b.Spaces[x] = make([]int, h)
	}
	return b
}

// Save ...
func (b *Board) Save(filename string) {

	// Open File
	file, err := os.Create(filename)
	dieOn(err)
	defer file.Close()

	// Write Dimensions
	fmt.Fprintf(file, "%d %d\n", b.W, b.H)

	// Write Starting
	fmt.Fprintf(file, "%d %d\n", b.StartX, b.StartY)

	for x := 0; x < b.W; x++ {
		for y := 0; y < b.H; y++ {
			switch {
			case b.Coin(x, y):
				fmt.Fprintf(file, "%d %d C\n", x, y)
			case b.Wall(x, y):
				fmt.Fprintf(file, "%d %d B\n", x, y)
			}
		}
	}
}

// Coin return true if the space contains a coin
func (b *Board) Coin(x, y int) bool {
	return b.valid(x, y) && (b.Spaces[x][y] == Coin)
}

// Wall return true if the space has a wall or if it is out of bounds
func (b *Board) Wall(x, y int) bool {
	return !b.valid(x, y) || (b.Spaces[x][y] == Wall)
}

// Object returns the object in the supplied space
func (b *Board) Object(x, y int) int {
	if !b.valid(x, y) {
		return None
	}
	return b.Spaces[x][y]
}

// Valid returns true if the supplied corrdinates are valid
func (b *Board) valid(x, y int) bool {

	// Bounds Check (i)
	if x < 0 || x >= b.W {
		return false
	}

	// Bounds Check (j)
	if y < 0 || y >= b.H {
		return false
	}

	// Valid!
	return true
}

// Add adds an object to the board
func (b *Board) Add(x, y, object int) {
	if !b.valid(x, y) {
		return
	}

	b.Spaces[x][y] = object
}

// Display ...
func (b *Board) Display() {

	// For Each Row (Downward)
	for y := b.H - 1; y >= 0; y-- {

		// For Each Column
		for x := 0; x < b.W; x++ {

			// Object
			object := b.Spaces[x][y]

			// Render Object
			switch {
			case object == Coin:
				fmt.Printf("o")
			case object == Wall:
				fmt.Printf("#")
			default:
				fmt.Printf("-")
			}
		}

		// Next Row
		fmt.Println()
	}
}
