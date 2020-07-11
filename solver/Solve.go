package solver

import (
	s "n-puzzle/structures"
	"container/list"
	"fmt"
)

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func ManhattanHeuristic(ctx *s.SContext) int {
	// fmt.Println("hello i am manhattan")
	score := 0
	for Y, line := range ctx.Puzzle {
		for X, nb := range line {
			score += abs(ctx.Final[nb].X - X) + abs(ctx.Final[nb].Y - Y)
		}
	}
	return score
}

func HeuristicPlacement(ctx *s.SContext) int {
	score := 0
	for Y, line := range  ctx.Puzzle {
		for X, nb := range line {
			if ctx.Final[nb].X != X || ctx.Final[nb].Y != Y {
				score += 1
			}
		}
	}
	return score
}

func CompleteHeuristic(ctx *s.SContext) int {
	for Y, line := range  ctx.Puzzle {
		for X, nb := range line {
			if ctx.Final[nb].X != X || ctx.Final[nb].Y != Y {
				return 1
			}
		}
	}
	return 0
}

func Solve(ctx *s.SContext) {
	node := s.SNode{Heuristic: ctx.Heuristic, Final: ctx.Final, Puzzle : ctx.Puzzle}
	lst := list.New()
	lst.PushFront(node)
	for e := lst.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
	return
}