package main

import (
	"fmt"
	"math/rand"
)

// Cell ...
type Cell struct {
	Char string
	Hit  bool
}

type Word struct {
	Text  string
	Found bool
}

// Board ...
type Board struct {
	R     int
	C     int
	Cells [][]*Cell
	Words []*Word
}

func (b Board) Draw() {

	// Dimensions
	fmt.Printf("%d %d\n", b.R, b.C)

	// Cells
	for r := 0; r < b.R; r++ {
		for c := 0; c < b.C; c++ {
			cell := b.Cells[r][c]
			fmt.Printf("%s", string(cell.Char))
		}
		fmt.Println()
	}

	// Words
	for _, w := range b.Words {
		fmt.Println(w.Text)
	}
}

// String ...
func (b Board) String() {

	// Cells
	for r := 0; r < b.R; r++ {
		for c := 0; c < b.C; c++ {
			cell := b.Cells[r][c]
			if !cell.Hit {
				fmt.Printf("%s", string(cell.Char))
			} else {
				fmt.Printf("\x1B[33m%s\033[0m", string(cell.Char))
			}
			fmt.Printf(" ")
		}
		fmt.Println()
	}

	// Words
	for _, w := range b.Words {
		if !w.Found {
			fmt.Println(w.Text)
		} else {
			fmt.Printf("\x1B[34m%s\033[0m\n", w.Text)
		}
	}
}

// NewBoard ...
func NewBoard(rows, cols int) *Board {
	b := &Board{R: rows, C: cols}
	b.Cells = make([][]*Cell, rows)
	for r := 0; r < rows; r++ {
		b.Cells[r] = make([]*Cell, cols)
		for c := 0; c < cols; c++ {
			b.Cells[r][c] = &Cell{Char: "."}
		}
	}
	return b
}

// Char returns the character at coordinate (x, y) on the board
func (b *Board) Char(x, y int) string {
	if x < 0 || x >= b.R {
		return ""
	}
	if y < 0 || y >= b.C {
		return ""
	}
	return b.Cells[x][y].Char
}

// Check checks if the supplied word fits the line-segment
func (b *Board) Check(r, c, dX, dY int, word string) bool {

	// Zero Delta (Nothing To Do)
	if dX == 0 && dY == 0 {
		return false
	}

	// For Each Character
	x := r
	y := c
	for z := 0; z < len(word); z++ {

		// Off Board!
		if b.Char(x, y) != string(word[z]) {
			return false
		}

		x += dX
		y += dY
	}

	x = r
	y = c
	for z := 0; z < len(word); z++ {
		b.Cells[x][y].Hit = true

		x += dX
		y += dY
	}

	return true
}

// Scan scans for a word
func (b *Board) Scan(r, c, dX, dY int) (string, bool) {

	// Zero Delta (Nothing To Do)
	if dX == 0 && dY == 0 {
		return "", false
	}

	// Length: 5 to 10
	for s := 10; s >= 6; s-- {

		// Build Word
		var word string

		x := r
		y := c
		var dead bool
		for z := 0; z < s && !dead; z++ {
			
			// Off Board!
			character := b.Char(x, y)
			if len(character) == 0 {
				dead = true
				break
			}

			word += character
			x += dX
			y += dY
		}

		// Check Word
		if !dead {
			if Valid(word) {
				return word, true
			}
		}
	}

	return "", false
}

// Solve will solve the board
func (b *Board) Solve() {
	for r := 0; r < b.R; r++ {
		for c := 0; c < b.C; c++ {
			for _, w := range b.Words {
				for dX := -1; dX <= 1; dX++ {
					for dY := -1; dY <= 1; dY++ {
						if b.Check(r, c, dX, dY, w.Text) {
							w.Found = true
						}
					}
				}
			}
		}
	}
}

func (b *Board) Random() {
	for r := 0; r < b.R; r++ {
		for c := 0; c < b.C; c++ {
			i := byte(65 + rand.Int() % 26)
			b.Cells[r][c].Char = string(i)
		}
	}

	for r := 0; r < b.R; r++ {
		for c := 0; c < b.C; c++ {
			for dX := -1; dX <= 1; dX++ {
				for dY := -1; dY <= 1; dY++ {
					if word, found := b.Scan(r, c, dX, dY); found {
						b.Words = append(b.Words, &Word{Text: word})
					}
				}
			}
		}
	}
}
