package solver

import (
	// l "container/list"
	// "fmt"
	s "n-puzzle/structures"
)

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func ManhattanHeuristic(puzzle [][]s.Tnumber, final []s.SVertex) int {
	// fmt.Println("hello i am manhattan")
	score := 0
	for Y, line := range puzzle {
		for X, nb := range line {
			score += Abs(int(final[nb].X)-X) + Abs(int(final[nb].Y)-Y)
		}
	}
	return score
}

func HeuristicPlacement(puzzle [][]s.Tnumber, final []s.SVertex) int {
	score := 0
	for Y, line := range puzzle {
		for X, nb := range line {
			if int(final[nb].X) != X || int(final[nb].Y) != Y {
				score++
			}
		}
	}
	return score
}

func CompleteHeuristic(puzzle [][]s.Tnumber, final []s.SVertex) int {
	for Y, line := range puzzle {
		for X, nb := range line {
			if int(final[nb].X) != X || int(final[nb].Y) != Y {
				return 1
			}
		}
	}
	return 0
}