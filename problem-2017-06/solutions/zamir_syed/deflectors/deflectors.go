package main

import (
	"fmt"
	"os"
)

func main() {

	grid := &Grid{}

	// Args
	switch {
	case len(os.Args) == 2:
		grid.Load(os.Args[1])
	case len(os.Args) == 5:
		rows := atoi(os.Args[1])
		cols := atoi(os.Args[2])
		dens := atof(os.Args[3])
		seed := atoi(os.Args[4])
		grid.Random(rows, cols, dens, seed)
	default:
		fmt.Fprintf(os.Stderr, "Two Versions:\n")
		fmt.Fprintf(os.Stderr, "-------------\n")
		fmt.Fprintf(os.Stderr, "deflectors GridFile\n")
		fmt.Fprintf(os.Stderr, "deflectors Rows Columns Density(0<x<1) Seed\n")
		os.Exit(1)
	}

	grid.Display()
	grid.Solve()
}
