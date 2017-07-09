package main

import (
	"fmt"
)

// Optimal ...
type Optimal struct {
	valid bool
	score int
	moves []int
}

// Smart ...
func Smart(b *Board) {

	// Copy Originals
	origX := b.StartX
	origY := b.StartY

	// Create Solution Grid
	optimals := make([][]*Optimal, b.W)
	for x := 0; x < b.W; x++ {
		optimals[x] = make([]*Optimal, b.H)
	}

	// Backward Induction (Northeast Corner)
	for x := b.W - 1; x >= 0; x-- {

		// North To South
		for y := b.H - 1; y >= 0; y-- {

			// Adjust Start To Current Grid-Position
			b.StartX = x
			b.StartY = y

			// Wall
			if b.Wall(x, y) {
				optimals[x][y] = &Optimal{valid: false}
				continue
			}

			////////////////////
			// Rightmost Edge //
			////////////////////

			if x == b.W-1 {
				player := NewPlayer(b)
				player.eat()
				player.fall()
				optimals[x][y] = &Optimal{valid: true, score: player.Score()}
				continue
			}

			////////////////////////
			// Measure: Jump-Walk //
			////////////////////////

			var scoreW int
			var movesW []int

			p := NewPlayer(b)
			p.eat()
			p.fall()

			// Will Jumping Help? (It Won't Hurt)
			jumpingHelped := false
			preJumpScore := p.Score()
			p.Jump()
			if p.Score() > preJumpScore {
				jumpingHelped = true
			}

			// Walk To Right
			if !p.wall(1, 0) {
				p.X++

				// Induction
				next := optimals[p.X][p.Y]
				if next.valid {
					scoreW += p.Score()
					scoreW += next.score

					if jumpingHelped {
						movesW = append(movesW, Jump)
					}

					movesW = append(movesW, Walk)
					movesW = append(movesW, next.moves...)
				}
			}

			////////////////////////
			// Measure: Jump-High //
			////////////////////////

			var scoreH int
			var movesH []int

			p = NewPlayer(b)
			p.eat()
			p.fall()

			// Will Jumping Help? (It Won't Hurt)
			jumpingHelped = false
			preJumpScore = p.Score()
			p.Jump()
			if p.Score() > preJumpScore {
				jumpingHelped = true
			}

			// Move North
			if !p.wall(0, 1) {
				p.Y++
				p.eat()

				// Move North
				if !p.wall(0, 1) {
					p.Y++
					p.eat()

					// Move Northeast
					if !p.wall(1, 1) {
						p.X++
						p.Y++

						// Induction
						next := optimals[p.X][p.Y]
						if next.valid {
							scoreH += p.Score()
							scoreH += next.score

							if jumpingHelped {
								movesH = append(movesH, Jump)
							}

							movesH = append(movesH, High)
							movesH = append(movesH, next.moves...)
						}
					}
				}
			}

			////////////////////////
			// Measure: Jump-Long //
			////////////////////////

			var scoreL int
			var movesL []int

			p = NewPlayer(b)
			p.eat()
			p.fall()

			// Will Jumping Help? (It Won't Hurt)
			jumpingHelped = false
			preJumpScore = p.Score()
			p.Jump()
			if p.Score() > preJumpScore {
				jumpingHelped = true
			}

			// Track Eastward Motion
			origX := p.X

			// Move Northeast
			if !p.wall(1, 1) {
				p.X++
				p.Y++

				// Move East (Eat Previous)
				if !p.wall(1, 0) {
					p.eat()
					p.X++

					// Move East (Eat Previous)
					if !p.wall(1, 0) {
						p.eat()
						p.X++
					}
				}
			}

			if p.X > origX {
				next := optimals[p.X][p.Y]
				if next.valid {
					scoreL += p.Score()
					scoreL += next.score

					if jumpingHelped {
						movesL = append(movesL, Jump)
					}

					movesL = append(movesL, Long)
					movesL = append(movesL, next.moves...)
				}
			}

			// Compute Optimal For Current Grid Position
			opt := &Optimal{}

			switch {

			// Winner: Jump-Walk
			case scoreW >= scoreH && scoreW >= scoreL && len(movesW) > 0:
				opt.score = scoreW
				opt.moves = movesW
				opt.valid = true

			// Winner: Jump-High
			case scoreH >= scoreW && scoreH >= scoreL && len(movesH) > 0:
				opt.score = scoreH
				opt.moves = movesH
				opt.valid = true

			// Winner: Jump-Long
			case scoreL >= scoreW && scoreL >= scoreH && len(movesL) > 0:
				opt.score = scoreL
				opt.moves = movesL
				opt.valid = true
			}

			optimals[x][y] = opt
		}
	}

	// Reset Original Starting
	b.StartX = origX
	b.StartY = origY

	// Best Solution
	bestMoves := optimals[b.StartX][b.StartY].moves
	winner := NewPlayer(b)

	// Draw
	fmt.Println("Best Solution")
	fmt.Println("-------------")
	winner.Play(bestMoves)
	winner.DrawIt()

	// No Solution!
	if len(bestMoves) == 0 {
		fmt.Println("Sorry - No Solution!!!")
		return
	}

	// Print Moves
	fmt.Println("Best-Moves:", Pretty(bestMoves))
	fmt.Println("Best-Coins:", winner.Score())
}
