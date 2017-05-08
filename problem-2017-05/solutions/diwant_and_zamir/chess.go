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
func solve(input string) {

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
}

func main() {

	var input string
	input += "BK h8\n"
	input += "WK f2\n"
	input += "WR b3\n"
	input += "WB e4\n"
	input += "WB h6\n"

	solve(input)
}
