package main

import (
	"fmt"
	"os"
)

func main() {
	var board *Board

	// Args
	switch {
	case len(os.Args) == 2:
		board = Build(os.Args[1])
	case len(os.Args) == 6:
		cols := atoi(os.Args[1])
		rows := atoi(os.Args[2])
		dens := atof(os.Args[3])
		seed := atoi(os.Args[4])
		save := os.Args[5]
		board = Random(cols, rows, dens, seed)
		board.Save(save)
	default:
		fmt.Fprintf(os.Stderr, "Two Versions:\n")
		fmt.Fprintf(os.Stderr, "-------------\n")
		fmt.Fprintf(os.Stderr, "mario GridFile\n")
		fmt.Fprintf(os.Stderr, "mario Width Height Density(0<x<1) Seed Save\n")
		os.Exit(1)
	}

	// Board
	NewPlayer(board).DrawIt()
	Solve(board)
}
