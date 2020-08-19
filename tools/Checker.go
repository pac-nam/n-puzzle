package tools

import (
	"fmt"
	s "n-puzzle/structures"
)

func findNumber(puzzle s.Tpuzzle, search s.Tnumber) s.SVertex {
	for Y, line := range puzzle {
		for X, nb := range line {
			if nb == search {
				return s.SVertex{X: s.Tnumber(X), Y: s.Tnumber(Y)}
			}
		}
	}
	return s.SVertex{X: 0, Y: 0}
}

func Checker(ctx *s.SContext, path []s.Tnumber) {
	fmt.Println("-------------------------Checker-------------------------")
	fmt.Println(ctx.Puzzle)
	fmt.Println("Sequence: ", path)
	copyPuzzle := CopyPuzzle(ctx.Puzzle, ctx.NSize)
	zero := findNumber(copyPuzzle, 0)
	for _, toSwap := range path {
		swapPlace := findNumber(copyPuzzle, toSwap)
		if (zero.X == swapPlace.X && zero.Y - 1 == swapPlace.Y) ||
			(zero.X == swapPlace.X && zero.Y + 1 == swapPlace.Y) ||
			(zero.X - 1 == swapPlace.X && zero.Y == swapPlace.Y) ||
			(zero.X + 1 == swapPlace.X && zero.Y == swapPlace.Y) {
				copyPuzzle[zero.Y][zero.X], copyPuzzle[swapPlace.Y][swapPlace.X] = copyPuzzle[swapPlace.Y][swapPlace.X], copyPuzzle[zero.Y][zero.X]
				zero = swapPlace
		} else {
			fmt.Println("wrong move trying to swap", toSwap, "in the puzzle:")
			fmt.Println(copyPuzzle)
		}
	}
	if ctx.Heuristic(copyPuzzle, ctx.Final) == 0 {
		fmt.Println("Check success")
	} else {
		fmt.Println("Check error:\n", copyPuzzle)
	}
	return
}