package solver

import (
    s "n-puzzle/structures"
    // "fmt"
)

func CopyPuzzle(puzzle [][]s.Tnumber, Size s.Tnumber) [][]s.Tnumber {
    newPuzzle := make([][]s.Tnumber, Size)
    for Y, line := range puzzle {
        newPuzzle[Y] = make([]s.Tnumber, Size)
        for X, nb := range line {
            newPuzzle[Y][X] = nb
        }
    }
    return newPuzzle
}

func PathCopy(path []s.Tnumber) []s.Tnumber {
	tmp := make([]s.Tnumber, len(path))
	copy(tmp, path)
	// fmt.Println("old: ", path, "\nnew: ", tmp, "\n")
	return tmp
}

func CopyNode(node *s.SNode) *s.SNode {
    tmp := make([]s.Tnumber, len(node.Path))
    copy(tmp, node.Path)
	return &s.SNode {
		Ctx:		node.Ctx,
		Cost:		node.Cost,
		Zero:		s.SVertex{X: node.Zero.X, Y: node.Zero.Y},
		Puzzle:		CopyPuzzle(node.Ctx.Puzzle, node.Ctx.NSize),
		Path:		PathCopy(tmp),
	}
}