package main

import (
	"fmt"
	// s "n-puzzle/structures"
	p "n-puzzle/parse"
	solve "n-puzzle/solver"
	t "n-puzzle/tools"
)

func main() {
	ctx, err := p.Parser()
	if err != "" {
		fmt.Println("error: " + err)
		return
	}
	// ctx.Heuristic(ctx)
	// fmt.Println(solver.HeuristicPlacement(ctx))
	// fmt.Println(ctx, "Normal end")
	res := solve.Astar(ctx)
	fmt.Println(res)
	t.Checker(ctx, res.Sequence)
}
