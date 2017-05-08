package main

// Space ...
type Space struct {
	row int
	col int
}

// IsOn ...
func (s *Space) isOn(row, col int) bool {
	return s.row == row && s.col == col
}

// IsOn ...
func (s *Space) Equals(t *Space) bool {
	return s.row == t.row && s.col == t.col
}

// Add ...
func (s *Space) Add(r, c int) {
	s.row += r
	s.col += c
}

func (s *Space) Valid() bool {
	switch {
	case s.row < 0:
		return false
	case s.col < 0:
		return false
	case s.row > 7:
		return false
	case s.col > 7:
		return false
	}
	return true
}

// King ...
type King struct {
	Space
}

// Rook ...
type Rook struct {
	Space
}

// Bishop ...
type Bishop struct {
	Space
}
