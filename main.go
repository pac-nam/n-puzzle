package main

import (
	"fmt"
	// s "n-puzzle/structures"
	p "n-puzzle/parse"
	// "n-puzzle/solver"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println(InvalidArgNumber)
		return
	}
	ctx, err := p.ParseFile(os.Args[1])
	if err != "" {
		fmt.Println("error: " + err)
		return
	}
	// fmt.Println(solver.HeuristicPlacement(ctx))
	fmt.Println(ctx, "Normal end")
}
