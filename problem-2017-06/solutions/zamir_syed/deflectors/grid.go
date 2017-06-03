package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
)

// Deflector ...
type Deflector struct {
	row  int
	col  int
	forw bool
}

// Grid ...
type Grid struct {
	rows int
	cols int
	tRow int
	tCol int
	defs []*Deflector
}

// Random creates a random grid using the supplied size and density
func (g *Grid) Random(rows int, cols int, density float64, seed int) {

	// Seed
	rand.Seed(int64(seed))

	// Set Dimensions
	g.rows = rows
	g.cols = cols

	// Generate Random Target
	g.tRow = rand.Intn(rows) + 1
	g.tCol = rand.Intn(cols) + 1

	// Add Deflectors
	for r := 1; r <= rows; r++ {
		for c := 1; c <= cols; c++ {

			// Target!
			if g.Hit(r, c) {
				continue
			}
			
			if rand.Intn(1000) < int(density * 1000) {
				g.AddDef(r, c, rand.Int() % 2 == 0)
			}
		}
	}
}

// Load loads a grid from the supplied filename
func (g *Grid) Load(filename string) {

	// Open File
	file, err := os.Open(filename)
	dieOn(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	
	// Read Dimensions
	scanner.Scan()
	line := scanner.Text()
	fmt.Sscanf(line, "%d %d", &g.rows, &g.cols)

	// Read Target
	scanner.Scan()
	line = scanner.Text()
	fmt.Sscanf(line, "%d %d", &g.tRow, &g.tCol)

	// Read Deflectors
	for scanner.Scan() {
		line := scanner.Text()
		var row int
		var col int
		var dir byte
		fmt.Sscanf(line, "%d %d %c", &row, &col, &dir)
		g.AddDef(row, col, dir == '/')
	}
}

// AddDef adds a deflector
func (g *Grid) AddDef(row int, col int, forw bool) {
	d := Deflector{row: row, col: col, forw: forw}
	g.defs = append(g.defs, &d)
}

// Hit returns true if the target is at the supplied position
func (g *Grid) Hit(row int, col int) bool {
	return (row == g.tRow) && (col == g.tCol)
}

// GetDef returns the deflector at the supplied position
func (g *Grid) GetDef(row int, col int) (*Deflector, bool) {
	for _, d := range g.defs {
		if d.row == row && d.col == col {
			return d, true
		}
	}
	return nil, false
}

// Display ...
func (g *Grid) Display() {

	// Rows
	for r := 1; r <= g.rows; r++ {

		// Begin Row
		fmt.Print("|")

		// Columns
		for c := 1; c <= g.cols; c++ {

			// Deflector
			def, ok := g.GetDef(r, c)
			if ok {
				if def.forw {
					fmt.Print("/")
				} else {
					fmt.Print("\\")
				}
				continue
			}
			
			// Target
			if g.Hit(r, c) {
				fmt.Print("X")
				continue
			}

			// Nothing
			fmt.Print(".")
		}

		// End Row
		fmt.Println("|")
	}
}

// Solve ...
func (g *Grid) Solve() {

	// Fire From West
	for r := 1; r <= g.rows; r++ {
		laser := &Laser{r, 0, East, g}
		if laser.Fire() {
			fmt.Println("W", r)
		}
	}

	// Fire From East
	for r := 1; r <= g.rows; r++ {
		laser := &Laser{r, g.cols+1, West, g}
		if laser.Fire() {
			fmt.Println("E", r)
		}
	}

	// Fire From North
	for c := 1; c <= g.cols; c++ {
		laser := &Laser{0, c, South, g}
		if laser.Fire() {
			fmt.Println("N", c)
		}
	}

	// Fire From South
	for c := 1; c <= g.cols; c++ {
		laser := &Laser{g.rows+1, c, North, g}
		if laser.Fire() {
			fmt.Println("S", c)
		}
	}
}
