package solver

import (
	n "n-puzzle/nstruct"
)

func heuristicPlacement(ctx *n.SContext) int {
	grid := ctx.Puzzle
	X, Xmin, Xmax, Y, Ymin, Ymax := 0, 0, ctx.NSize-1, 0, 1, ctx.NSize-1
	square := ctx.NSize * ctx.NSize
	score := 0
	for i := 1; i < square; {
		for X < Xmax {
			if grid[Y][X] != i {
				score++
			}
			X++
			i++
		}
		Xmax--
		for Y < Ymax {
			if grid[Y][X] != i {
				score++
			}
			Y++
			i++
		}
		Ymax--
		for X > Xmin {
			if grid[Y][X] != i {
				score++
			}
			X--
			i++
		}
		Xmin++
		for Y > Ymin {
			if grid[Y][X] != i {
				score++
			}
			Y--
			i++
		}
		Ymin++
	}

	return (score)
}
