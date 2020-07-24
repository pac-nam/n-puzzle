package solver

import (
	s "n-puzzle/structures"
	// "fmt"
)

func firstRoutine(ctx *s.SContext) {
	firstNode := s.SNode {
		Ctx:		ctx,
		Cost:		0,
		Zero:		ctx.Zero,
		Puzzle:		ctx.Puzzle,
		Path:		[]uint16{0},
	}
	go Algo(&firstNode)
}