package main

import (
	"fmt"
	// s "n-puzzle/structures"
	p "n-puzzle/parse"
	// "n-puzzle/solver"
)

func main() {
	ctx, err := p.Parser()
	if err != "" {
		fmt.Println("error: " + err)
		return
	}
	// ctx.Heuristic(ctx)
	// fmt.Println(solver.HeuristicPlacement(ctx))
	fmt.Println(ctx, "Normal end")
}
