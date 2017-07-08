package main

import (
	"fmt"
)

// Solve ...
func Solve(b *Board, startX int, startY int) {

	var moves []int
	var bestScore int
	var bestMoves []int

	// Ensure Zero Coins A Potential Winner
	bestScore = -1

	for {

		// New Player
		player := NewPlayer(b, startX, startY)
		ending := player.Play(moves)

		// New High-Score?
		if ending == Done && player.Score() > bestScore {
			fmt.Println(Pretty(moves), player.Score())

			// Record Best
			bestScore = player.Score()
			bestMoves = make([]int, len(moves))
			copy(bestMoves, moves)
		}

		// Increment
		if !Incr(&moves, ending) {
			break
		}
	}

	// Draw
	winner := NewPlayer(b, startX, startY)
	winner.Play(bestMoves)
	winner.DrawIt()

	// No Solution!
	if len(bestMoves) == 0 {
		fmt.Println("Sorry - No Solution!!!")
	}
}

// Pretty ...
func Pretty(moves []int) string {
	var result string
	for _, m := range moves {
		switch m {
		case Walk:
			result += "W"
		case Jump:
			result += "J"
		case High:
			result += "H"
		case Long:
			result += "L"
		default:
			result += "?"
		}
	}
	return result
}

// Incr ...
func Incr(moves *[]int, ending int) bool {
	m := len(*moves)

	switch {

	// Good Move: Go Again
	case ending == Good:
		(*moves) = append(*moves, Walk)
		return true

	// End Of The Line: Replace Previous Moves
	default:

		// Find (Backwards) Candidate-Move For Increment
		for i := m - 1; i >= 0; i-- {

			// Try Next Option
			cand := (*moves)[i]
			if cand < Long {
				(*moves)[i]++

				// Redundancy: Don't Jump Twice In A Row
				if i-1 >= 0 && (*moves)[i-1] == Jump && (*moves)[i] == Jump {
					(*moves)[i]++
				}

				// Redundancy: Don't High-Jump After A Jump
				if i-1 >= 0 && (*moves)[i-1] == Jump && (*moves)[i] == High {
					(*moves)[i]++
				}

				// Truncate To Current Position
				(*moves) = (*moves)[:i+1]
				return true
			}
		}
	}

	return false
}
