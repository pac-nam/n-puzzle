package solver

import (
    s "n-puzzle/structures"
    // "fmt"
)

func CopyPuzzle(puzzle [][]uint16, Size uint16) [][]uint16 {
    newPuzzle := make([][]uint16, Size)
    for Y, line := range puzzle {
        newPuzzle[Y] = make([]uint16, Size)
        for X, nb := range line {
            newPuzzle[Y][X] = nb
        }
    }
    return newPuzzle
}

func PathCopy(path []uint16) []uint16 {
	tmp := make([]uint16, len(path))
	copy(tmp, path)
	// fmt.Println("old: ", path, "\nnew: ", tmp, "\n")
	return tmp
}

func CopyNode(node *s.SNode) *s.SNode {
    tmp := make([]uint16, len(node.Path))
    copy(tmp, node.Path)
	return &s.SNode {
		Ctx:		node.Ctx,
		Cost:		node.Cost,
		Zero:		s.SVertex{X: node.Zero.X, Y: node.Zero.Y},
		Puzzle:		CopyPuzzle(node.Ctx.Puzzle, node.Ctx.NSize),
		Path:		PathCopy(tmp),
	}
}