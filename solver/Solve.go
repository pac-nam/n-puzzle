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


func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func littleManhattan(nb, Ystart, Xstart, NSize int) int {
	X, Xmin, Xmax, Y, Ymin, Ymax := 0, 0, NSize-1, 0, 1, NSize-1
	square := NSize * NSize
	for i := 1; i < square; {
		for X < Xmax {
			if nb == i {
				return abs(X-Xstart) + abs(Y-Ystart)
			}
			i++
			X++
		}
		Xmax--
		for Y < Ymax {
			if nb == i {
				return abs(X-Xstart) + abs(Y-Ystart)
			}
			i++
			Y++
		}
		Ymax--
		for X > Xmin {
			if nb == i {
				return abs(X-Xstart) + abs(Y-Ystart)
			}
			i++
			X--
		}
		Xmin++
		for Y > Ymin {
			if nb == i {
				return abs(X-Xstart) + abs(Y-Ystart)
			}
			i++
			Y-- 
		}
		Ymin++
	}
	return 0
}

func manhattanHeuristic(ctx *n.SContext) int {
	score := 0
	for Y, line := range ctx.Puzzle {
		for X, nb := range line {
			score += littleManhattan(nb, Y, X, ctx.NSize)
		}
	}
	return score
}