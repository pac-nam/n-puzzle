package parse

import (
	n "n-puzzle/nstruct"
	"fmt"
)

func checkPuzzle(ctx *n.SContext) (*n.SContext, string) {
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
	fmt.Println(check)
	return ctx, ""
}