package heuristic

import (
	s "n-puzzle/structures"
	"math"
)

func Euclidian(puzzle [][]s.Tnumber, final []s.SVertex) int {
	var score float64 = 0
	for Y, line := range puzzle {
		for X, nb := range line {
			score += math.Abs(float64((int(final[nb].X)-X))) + math.Abs(float64((int(final[nb].Y)-Y)))
		}
	}
	return int(score)
}