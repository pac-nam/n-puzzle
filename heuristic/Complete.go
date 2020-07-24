package heuristic

import (
	s "n-puzzle/structures"
)

func Complete(puzzle [][]s.Tnumber, final []s.SVertex) int {
	for Y, line := range puzzle {
		for X, nb := range line {
			if int(final[nb].X) != X || int(final[nb].Y) != Y {
				return 1
			}
		}
	}
	return 0
}