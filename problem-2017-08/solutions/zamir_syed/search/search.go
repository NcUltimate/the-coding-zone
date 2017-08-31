package main

import (
	"os"
)

func main() {

	if len(os.Args) == 1 {
		LoadWords("english-words/words.txt")
		board := NewBoard(1000, 1000)
		board.Random()
		board.Draw()
	} else {
		filename := os.Args[1]
		board := Build(filename)
		board.Solve()
		board.String()
	}
}
