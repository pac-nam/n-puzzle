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

func ManhattanHeuristic(puzzle [][]int, final []s.SVertex) int {
	// fmt.Println("hello i am manhattan")
	score := 0
	for Y, line := range puzzle {
		for X, nb := range line {
			score += abs(final[nb].X - X) + abs(final[nb].Y - Y)
		}
	}
	return score
}

func HeuristicPlacement(puzzle [][]int, final []s.SVertex) int {
	score := 0
	for Y, line := range  puzzle {
		for X, nb := range line {
			if final[nb].X != X || final[nb].Y != Y {
				score ++
			}
		}
	}
	return score
}

func CompleteHeuristic(puzzle [][]int, final []s.SVertex) int {
	for Y, line := range  puzzle {
		for X, nb := range line {
			if final[nb].X != X || final[nb].Y != Y {
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