package parse

import (
	s "n-puzzle/structures"
	"fmt"
)
func createFinal(ctx *s.SContext) []s.SVertex {
	X, Xmin, Xmax, Y, Ymin, Ymax := 0, 0, ctx.NSize-1, 0, 1, ctx.NSize-1
	square := ctx.NSize * ctx.NSize
	final := make([]s.SVertex, square)
	for i := 1; i < square; {
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
	check := make([]bool, ctx.NSize * ctx.NSize)
	max := ctx.NSize * ctx.NSize - 1
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
	return ctx, ""
}