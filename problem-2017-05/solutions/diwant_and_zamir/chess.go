package main

import (
	"fmt"
	"strings"
)

// Board ...
type Board struct {
	rooks   []*Rook
	bishops []*Bishop
	wKing   *King
	bKing   *King
}

func (b *Board) Occupied(s Space) bool {

	if b.wKing.isOn(s.row, s.col) {
		return true
	}

	if b.bKing.isOn(s.row, s.col) {
		return true
	}

	for _, rook := range b.rooks {
		if rook.isOn(s.row, s.col) {
			return true
		}
	}

	for _, bishop := range b.bishops {
		if bishop.isOn(s.row, s.col) {
			return true
		}
	}

	return false
}

func (b *Board) String() string {
	disp := make([][]rune, 8)

	// Fill With Spaces
	for i := 0; i < 8; i++ {
		disp[i] = make([]rune, 8)
	}

	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			disp[i][j] = '.'
		}
	}

	// Add Kings
	disp[b.bKing.col][b.bKing.row] = 'B'
	disp[b.wKing.col][b.wKing.row] = 'W'

	for _, rook := range b.rooks {
		disp[rook.col][rook.row] = 'r'
	}

	for _, bishop := range b.bishops {
		disp[bishop.col][bishop.row] = 'b'
	}

	var result string
	for i := 7; i >= 0; i-- {
		for j := 0; j <= 7; j++ {
			result += string(disp[i][j])
		}
		result += "\n"
	}
	return result
}

// Solve
func solve(input string) []*Board {
	fmt.Println("---------------")
	fmt.Println("-- New Board --")
	fmt.Println("---------------")
	fmt.Println("")
	b := &Board{}

	// Lines
	lines := strings.Split(input, "\n")

	// For Each Line
	for _, line := range lines {

		if len(line) != 5 {
			continue
		}

		// Position
		row := int(line[3]) - 97
		col := int(line[4]) - 49

		switch {
		case line[0] == 'B':
			b.bKing = &King{}
			b.bKing.row = row
			b.bKing.col = col

		case line[0] == 'W' && line[1] == 'K':
			b.wKing = &King{}
			b.wKing.row = row
			b.wKing.col = col

		case line[0] == 'W' && line[1] == 'R':
			rook := &Rook{}
			rook.row = row
			rook.col = col
			b.rooks = append(b.rooks, rook)

		case line[0] == 'W' && line[1] == 'B':
			bishop := &Bishop{}
			bishop.row = row
			bishop.col = col
			b.bishops = append(b.bishops, bishop)
		}
	}

	fmt.Println(b.String())
	fmt.Println()
	fmt.Println("Check:", b.check())
	fmt.Println("Check Mate:", b.checkMate())

	// Store solutions here
	solutions := make([]*Board, 0)

	// Check Rooks
	for _, rook := range b.rooks {

		// Rook Directions
		deltas := []Space{
			Space{0, 1},
			Space{0, -1},
			Space{1, 0},
			Space{-1, 0},
		}

		for _, delta := range deltas {
			blocked := false
			for i := 1; i <= 8 && !blocked; i++ {

				// Check if new space is occupied
				space := rook.Space
				space.Add(i*delta.row, i*delta.col)
				if b.Occupied(space) {
					blocked = true
					continue
				}

				// Move rook and check for mate
				rook.Add(i*delta.row, i*delta.col)
				if rook.Valid() && b.checkMate() {
					solutions = append(solutions, b.copy())
				}

				// Restore Rook
				rook.Add(-i*delta.row, -i*delta.col)
			}
		}
	}

	// Check Bishops
	for _, bishop := range b.bishops {

		// Bishop Directions
		deltas := []Space{
			Space{1, 1},
			Space{1, -1},
			Space{-1, 1},
			Space{-1, -1},
		}

		for _, delta := range deltas {
			blocked := false
			for i := 1; i <= 8 && !blocked; i++ {

				// Check if new space is occupied
				space := bishop.Space
				space.Add(i*delta.row, i*delta.col)
				if b.Occupied(space) {
					blocked = true
					continue
				}

				// Move bishop and check for mate
				bishop.Add(i*delta.row, i*delta.col)
				if bishop.Valid() && b.checkMate() {
					solutions = append(solutions, b.copy())
				}

				// Restore Bishop
				bishop.Add(-i*delta.row, -i*delta.col)
			}
		}
	}
	return solutions
}

func main() {

	var input string
	input += "BK h8\n"
	input += "WK f2\n"
	input += "WR b3\n"
	input += "WB e4\n"
	input += "WB h6\n"

	solutions := solve(input)
	for _, s := range solutions {
		fmt.Println(s)
	}

	// Rook Mate
	input = ""
	input += "BK f8\n"
	input += "WK f6\n"
	input += "WR b3\n"

	solve(input)
	solutions = solve(input)
	for _, s := range solutions {
		fmt.Println(s)
	}

	input = `BK h8
WB f8
WB e7
WB f7
WB g6
WB g5
WB h4
WK a1
`
	solve(input)
	solutions = solve(input)
	for _, s := range solutions {
		fmt.Println(s)
	}
}
