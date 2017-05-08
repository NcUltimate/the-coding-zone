package main

func (b *Board) check() bool {

	// Are Rooks Checking Black-King?
	for _, rook := range b.rooks {
		if b.rookCheck(rook.Space) {
			return true
		}
	}

	// Are Bishops Checking Black-King?
	for _, bishop := range b.bishops {
		if b.bishopCheck(bishop.Space) {
			return true
		}
	}

	return false
}

func (b *Board) rookCheck(space Space) bool {

	// Directions
	deltas := []Space{
		Space{0, 1},
		Space{0, -1},
		Space{1, 0},
		Space{-1, 0},
	}

	// For Each Direction
	for _, delta := range deltas {
		if b.deltaCheck(space, delta) {
			return true
		}
	}

	return false
}

func (b *Board) bishopCheck(space Space) bool {

	// Directions
	deltas := []Space{
		Space{1, 1},
		Space{1, -1},
		Space{-1, 1},
		Space{-1, -1},
	}

	// For Each Direction
	for _, delta := range deltas {
		if b.deltaCheck(space, delta) {
			return true
		}
	}

	return false
}

func (b *Board) deltaCheck(piece Space, delta Space) bool {

	// Set Space To Rook
	space := Space{row: piece.row, col: piece.col}

	// Add Delta
	space.Add(delta.row, delta.col)

	// While Still On Board
	for space.Valid() {

		// Checking Black-King
		if space.Equals(&b.bKing.Space) {
			return true
		}

		// Obstructed
		if b.Occupied(space) {
			return false
		}

		// Next Space
		space.Add(delta.row, delta.col)
	}

	return false
}
