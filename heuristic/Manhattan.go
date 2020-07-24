package heuristic

import (
	s "n-puzzle/structures"
	t "n-puzzle/tools"
)

func Manhattan(puzzle [][]s.Tnumber, final []s.SVertex) int {
	score := 0
	for Y, line := range puzzle {
		for X, nb := range line {
			score += t.Abs(int(final[nb].X)-X) + t.Abs(int(final[nb].Y)-Y)
		}
	}
	return score
}