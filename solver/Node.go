package solver

import (
	s "n-puzzle/structures"
)

func NewNode(ctx s.SContext) s.SNode {
	return s.SNode{
        Heuristic: ctx.Heuristic,
        Final: ctx.Final,
        Puzzle: NodeCopyPuzzle(ctx.Puzzle, ctx.NSize),
        Zero: s.SVertex{X : ctx.Zero.X, Y : ctx.Zero.Y},
		Size: ctx.NSize,
	}
}

func NewNextNode(puzzle [][]int, zero s.SVertex, node s.SNode) s.SNode {
	return s.SNode{
		Heuristic : node.Heuristic,
		Final : node.Final,
		Puzzle : NodeCopyPuzzle(puzzle, node.Size),
		Zero : s.SVertex{X : zero.X, Y : zero.Y},
		Size : node.Size,
	}
}

func NodeCopyPuzzle(puzzle [][]int, Size int) [][]int {
    newPuzzle := make([][]int, Size)
    for Y, line := range puzzle {
        newPuzzle[Y] = make([]int, Size)
        for X, nb := range line {
            newPuzzle[Y][X] = nb
        }
    }
    return newPuzzle
}