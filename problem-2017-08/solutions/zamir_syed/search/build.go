package main

import (
	"fmt"
	"os"
)

// Build returns a new board and the starting coordinates
func Build(filename string) *Board {

	// Open File
	file, err := os.Open(filename)
	dieOn(err)

	// Read Dimensions
	var rows int
	var cols int
	fmt.Fscanf(file, "%d %d\n", &rows, &cols)

	// New Board
	b := NewBoard(rows, cols)

	// Read Board
	for r := 0; r < rows; r++ {
		row := make([]byte, cols+1)
		file.Read(row)
		for c := 0; c < cols; c++ {
			b.Cells[r][c].Char = string(row[c])
		}
	}

	// Read Words
	for {
		var word string
		if n, _ := fmt.Fscanf(file, "%s\n", &word); n == 0 {
			break
		}
		w := &Word{Text: word}
		b.Words = append(b.Words, w)
	}

	return b
}
