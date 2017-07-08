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
	Coins map[int]bool
	Path  []Space
}

// NewPlayer ...
func NewPlayer(b *Board, x, y int) *Player {
	return &Player{
		B:     b,
		X:     x,
		Y:     y,
		Coins: make(map[int]bool),
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
		case m == High:
			success = p.High()
		case m == Long:
			success = p.Long()
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

// Long performs a long-jump (NE,E,E)
func (p *Player) Long() bool {

	// See If Moved Right By End
	origX := p.X

	// No Wall NE
	if !p.wall(1, 1) {
		p.X++
		p.Y++
		p.touch()

		// No Wall East?
		if !p.wall(1, 0) {
			p.X++
			p.touch()

			// No Wall East?
			if !p.wall(1, 0) {
				p.X++
				p.touch()
			}
		}
	}

	// Now Fall
	p.fall()

	// Success
	if p.X > origX {
		return true
	}

	// Not Useful
	return false
}

// High performs a high-jump (N,N,NE)
func (p *Player) High() bool {

	// See If Moved Right By End
	origX := p.X

	// No Wall Above?
	if !p.wall(0, 1) {
		p.Y++
		p.touch()

		// No Wall Above?
		if !p.wall(0, 1) {
			p.Y++
			p.touch()

			// Wall North-East?
			if !p.wall(1, 1) {
				p.X++
				p.Y++
				p.touch()
			}
		}
	}

	// Now Fall
	p.fall()

	// Success
	if p.X > origX {
		return true
	}

	// Not Useful
	return false
}

// Jump performs a jump (N,N)
func (p *Player) Jump() bool {

	// No Wall Above?
	if !p.wall(0, 1) {

		// Move Up One
		p.Y++
		p.touch()

		// No Wall Above?
		if !p.wall(0, 1) {

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
	if p.wall(1, 0) {
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

	// Compute ID Of Coin
	coinID := y*p.B.W + x

	// Check If Coin Already Eaten
	_, ok := p.Coins[coinID]
	if ok {
		return false
	}

	// Eat Coin!
	p.Coins[coinID] = true
	return true
}

// Fall falls until a wall is beneath
func (p *Player) fall() {

	// While No Wall Beneath Player
	for !p.wall(0, -1) {
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

// Score ...
func (p *Player) Score() int {
	return len(p.Coins)
}

// Wall returns true if there is a wall in the delta-position
func (p *Player) wall(dX, dY int) bool {
	return p.B.Wall(p.X+dX, p.Y+dY)
}

// Touched returns true of the space was touched
func (p *Player) touched(x, y int) bool {
	for _, s := range p.Path {
		if s.X == x && s.Y == y {
			return true
		}
	}
	return false
}

// DrawIt draws the board with touched moves!
func (p *Player) DrawIt() {

	// For Each Row (Downward)
	for y := p.B.H - 1; y >= 0; y-- {

		// For Each Column
		for x := 0; x < p.B.W; x++ {

			// Object
			object := p.B.Spaces[x][y]

			// Touched?
			touched := p.touched(x, y)

			// Render Object
			switch {
			case object == Coin && touched:
				fmt.Printf(green("o"))
			case object == Coin && !touched:
				fmt.Printf("o")
			case object == Wall:
				fmt.Printf(block(" "))
			case object == None && touched:
				fmt.Printf(blue("-"))
			default:
				fmt.Printf(".")
			}
		}

		// Next Row
		fmt.Println()
	}
}
