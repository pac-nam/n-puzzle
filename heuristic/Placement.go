package heuristic

import (
	s "n-puzzle/structures"
)

func Placement(puzzle [][]s.Tnumber, final []s.SVertex) int {
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