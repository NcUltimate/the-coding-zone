package main

import (
	"fmt"
)

// Moves ...
const (
	Walk = 0
	Jump = 1
	High = 2
	Long = 3
)

// Ending ...
const (
	Good = 0
	Nope = 1
	Done = 2
)

// Space ...
type Space struct {
	X int
	Y int
	C bool
}

// Player ...
type Player struct {
	B     *Board
	X     int
	Y     int
	Coins map[int]int
	Path  []Space
}

// NewPlayer ...
func NewPlayer(b *Board, x, y int) *Player {
	return &Player{
		B:     b,
		X:     x,
		Y:     y,
		Coins: make(map[int]int),
	}
}

// Play makes the supplied moves.
// It returns the ending-state of the player.
func (p *Player) Play(moves []int) int {

	// Initial Fall From Starting Position
	p.touch()
	p.fall()

	// For Each Move
	for _, m := range moves {

		var success bool

		// Make Move
		switch {
		case m == Walk:
			success = p.Walk()
		case m == Jump:
			success = p.Jump()
		}

		// Success?
		if !success {
			return Nope
		}
	}

	// Reached End!
	if p.X == p.B.W-1 {
		return Done
	}

	// Good Move But Not Done
	return Good
}

// Jump jumps up two spaces
func (p *Player) Jump() bool {

	// No Wall Above?
	if !p.B.Wall(p.X, p.Y+1) {

		// Move Up One
		p.Y++
		p.touch()

		// No Wall Above?
		if !p.B.Wall(p.X, p.Y+1) {

			// Move Up Again
			p.Y++
			p.touch()
		}
	}

	// Now Fall
	p.fall()
	return true
}

// Walk ...
func (p *Player) Walk() bool {

	// At Right Most Edge
	if p.X == p.B.W-1 {
		return false
	}

	// Wall To Right
	if p.B.Wall(p.X+1, p.Y) {
		return false
	}

	// Move To the Right
	p.X++
	p.touch()

	// Fall
	p.fall()

	return true
}

// Eat eats a coin in the position
func (p *Player) eat() bool {

	// Coordinates
	x := p.X
	y := p.Y

	// No Coin In Spot
	if !p.B.Coin(x, y) {
		return false
	}

	// Check If Coin Already Eaten
	coinY, ok := p.Coins[x]
	if ok || coinY == y {
		return false
	}

	// Eat Coin!
	p.Coins[x] = y
	return true
}

// Fall falls until a wall is beneath
func (p *Player) fall() {

	// While No Wall Beneath Player
	for !p.B.Wall(p.X, p.Y-1) {
		p.Y--
		p.touch()
	}
}

// Touch touches the current square by
// eating a coin (if there is one) and adding
// the position to the player's path.
func (p *Player) touch() {
	coin := p.eat()
	space := Space{p.X, p.Y, coin}
	p.Path = append(p.Path, space)
}

// Display
func (p *Player) Display() {
	for _, p := range p.Path {
		fmt.Printf("(%d,%d)", p.X, p.Y)
		if p.C {
			fmt.Printf(" Coin")
		}
		fmt.Println()
	}
}
