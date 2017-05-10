package main

func (b *Board) checkMate() bool {
	if !b.check() {
		return false
	}

	deltas := []Space{
		Space{0, 1},
		Space{0, -1},
		Space{1, 0},
		Space{-1, 0},
		Space{-1, -1},
		Space{-1, 1},
		Space{1, -1},
		Space{1, 1},
	}

	for _, delta := range deltas {
		inCheck := true

		// Move BK
		b.bKing.Add(delta.row, delta.col)

		// Check?
		if b.bKing.Valid() && !b.check() && !b.kingsTouch() {
			inCheck = false
		}

		// Restore BK
		b.bKing.Add(-delta.row, -delta.col)

		if !inCheck {
			return false
		}
	}

	return true
}

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

// kingsTouch gets the shortest distance between the two kings
func (b *Board) kingsTouch() bool {
	rowDist := absoluteInt(b.bKing.row - b.wKing.row)
	colDist := absoluteInt(b.bKing.col - b.wKing.col)

	return rowDist <= 1 && colDist <= 1
}

func absoluteInt(val int) int {
	if val < 0 {
		return -val
	}
	return val
}

func (b *Board) copy() *Board {
	r := Board{}

	// Deep Copy Bishops
	for _, bishop := range b.bishops {
		newBishop := *bishop
		r.bishops = append(r.bishops, &newBishop)
	}

	// Deep Copy Rooks
	for _, rook := range b.rooks {
		newRook := *rook
		r.rooks = append(r.rooks, &newRook)
	}

	newBKing := *b.bKing
	r.bKing = &newBKing

	newWKing := *b.wKing
	r.wKing = &newWKing

	return &r
}
