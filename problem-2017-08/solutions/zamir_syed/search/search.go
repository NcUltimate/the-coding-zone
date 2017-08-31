package main

import (
	"fmt"
	"os"
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

// String ...
func (b Board) String() string {
	var result string

	// Cells
	for r := 0; r < b.R; r++ {
		for c := 0; c < b.C; c++ {
			cell := b.Cells[r][c]
			if !cell.Hit {
				result += fmt.Sprintf("%s", string(cell.Char))
			} else {
				result += fmt.Sprintf("\x1B[33m%s\033[0m", string(cell.Char))
			}
		}
		result += "\n"
	}

	// Words
	result += "\n"
	for _, w := range b.Words {
		if !w.Found {
			result += w.Text + "\n"
		} else {
			result += fmt.Sprintf("\x1B[34m%s\033[0m\n", w.Text)
		}
	}
	return result
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

func main() {
	filename := os.Args[1]
	board := Build(filename)
	board.Solve()
	fmt.Printf(board.String())
}
