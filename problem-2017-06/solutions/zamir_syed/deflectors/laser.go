package main

// Directions
const (
	North = iota
	South
	East
	West
)

// Laser ...
type Laser struct {
	row  int
	col  int
	dir  int
	grid *Grid
}

// Fire moves the laser position until it reaches the target or escapes
func (l *Laser) Fire() bool {

	for {

		// Move One Unit
		l.move()

		// Hit Target?
		if l.hit() {
			return true
		}

		// Escaped
		if l.escaped() {
			return false
		}
	}
}

// Escape fires the laser hoping to escape the board
func (l *Laser) Escape() bool {

	for {

		// Move One Unit
		l.move()

		// Hit Target?
		if l.hit() {
			return false
		}

		// Escaped
		if l.escaped() {
			return true
		}
	}
}

// Move moves the laser one unit
func (l *Laser) move() {

	// Move One Unit
	switch {
	case l.dir == North:
		l.row--
	case l.dir == South:
		l.row++
	case l.dir == East:
		l.col++
	case l.dir == West:
		l.col--
	}

	// Deflect?
	def, present := l.grid.GetDef(l.row, l.col)
	if present {
		switch {
		case def.forw && l.dir == North:
			l.dir = East
		case def.forw && l.dir == South:
			l.dir = West
		case def.forw && l.dir == East:
			l.dir = North
		case def.forw && l.dir == West:
			l.dir = South
		case !def.forw && l.dir == North:
			l.dir = West
		case !def.forw && l.dir == South:
			l.dir = East
		case !def.forw && l.dir == East:
			l.dir = South
		case !def.forw && l.dir == West:
			l.dir = North
		}
	}
}

// Hit returns true if the laser has reached the target
func (l *Laser) hit() bool {
	return l.grid.Hit(l.row, l.col)
}

// Hit returns true if the laser has escaped the grid
func (l *Laser) escaped() bool {
	switch {
	case l.row < 1:
		return true
	case l.row > l.grid.rows:
		return true
	case l.col < 1:
		return true
	case l.col > l.grid.cols:
		return true
	default:
		return false
	}
}
