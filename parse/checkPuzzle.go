package parse

import (
	"fmt"
	solv "n-puzzle/solver"
	s "n-puzzle/structures"
)

func distHeuristic(puzzle [][]s.Tnumber, final []s.SVertex) int {
	x0, y0 := 0, 0
	for Y, line := range puzzle {
		for X, nb := range line {
			if nb == 0 {
				y0 = Y
				x0 = X
				break
			}
		}
	}
	return (solv.Abs(int(final[0].X)-x0) + solv.Abs(int(final[0].Y)-y0))
}

func checkSolved(ctx *s.SContext) bool {
	cPuzzle := solv.CopyPuzzle(ctx.Puzzle, ctx.NSize)
	dist := distHeuristic(cPuzzle, ctx.Final)
	nbSwap := 0
	pair := dist % 2
	if cPuzzle[ctx.Final[0].Y][ctx.Final[0].X] != cPuzzle[ctx.Zero.Y][ctx.Zero.X] {
		cPuzzle[ctx.Final[0].Y][ctx.Final[0].X], cPuzzle[ctx.Zero.Y][ctx.Zero.X] = cPuzzle[ctx.Zero.Y][ctx.Zero.X], cPuzzle[ctx.Final[0].Y][ctx.Final[0].X]
		nbSwap++
	}
	for nb := ctx.NSize*ctx.NSize - 1; nb > 0; nb-- {
		y0, x0 := 0, 0
		if cPuzzle[ctx.Final[nb].Y][ctx.Final[nb].X] != s.Tnumber(nb) {
			for Y, line := range cPuzzle {
				for X, i := range line {
					if i == s.Tnumber(nb) {
						y0 = Y
						x0 = X
						break
					}
				}
			}
			cPuzzle[ctx.Final[nb].Y][ctx.Final[nb].X], cPuzzle[y0][x0] = cPuzzle[y0][x0], cPuzzle[ctx.Final[nb].Y][ctx.Final[nb].X]
			nbSwap++
		}
	}
	if pair == nbSwap%2 {
		return true
	}
	return false
}

func createFinal(ctx *s.SContext) []s.SVertex {
	var X, Xmin, Xmax, Y, Ymin, Ymax s.Tnumber = 0, 0, ctx.NSize-1, 0, 1, ctx.NSize-1
	var square, i s.Tnumber = ctx.NSize * ctx.NSize, 1
	final := make([]s.SVertex, square)
	for i < square {
		for X < Xmax {
			final[i] = s.SVertex{X: X, Y: Y}
			X++
			i++
		}
		Xmax--
		for Y < Ymax {
			final[i] = s.SVertex{X: X, Y: Y}
			Y++
			i++
		}
		Ymax--
		for X > Xmin {
			final[i] = s.SVertex{X: X, Y: Y}
			X--
			i++
		}
		Xmin++
		for Y > Ymin {
			final[i] = s.SVertex{X: X, Y: Y}
			Y--
			i++
		}
		Ymin++
	}
	final[0] = s.SVertex{X: X, Y: Y}
	return final
}

func checkPuzzle(ctx *s.SContext) (*s.SContext, string) {
	check := make([]bool, ctx.NSize*ctx.NSize)
	max := s.Tnumber(ctx.NSize*ctx.NSize - 1)
	for _, line := range ctx.Puzzle {
		for _, nb := range line {
			if nb < 0 || nb > max {
				return ctx, "Invalid number " + fmt.Sprint(nb)
			} else if check[nb] {
				return ctx, "Multiple occurence of number " + fmt.Sprint(nb)
			}
			check[nb] = true
		}
	}
	ctx.Final = createFinal(ctx)
	solved := checkSolved(ctx)
	if solved == false {
		return ctx, "Unsolvable puzzle"
	}
	return ctx, ""
}
