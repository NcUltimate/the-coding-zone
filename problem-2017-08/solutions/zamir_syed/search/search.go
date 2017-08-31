package main

import (
	"fmt"
	"os"
)

func main() {
	filename := os.Args[1]
	board := Build(filename)
	board.Solve()
	fmt.Printf(board.String())
}
