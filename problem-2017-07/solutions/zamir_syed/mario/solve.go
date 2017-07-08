package main

import ()

// Solve ...
func Solve(b *Board, startX int, startY int) {
	player := NewPlayer(b, startX, startY)

	moves := []int{
		Long,
		Walk,
		High,
		Walk,
		Jump,
		Walk,
		Jump,
		Walk,
		Walk,
		Walk,
		Jump,
		Walk,
		Jump,
	}
	player.Play(moves)
	player.Display()
}
