package structures

import (
	"fmt"
)

type SNode struct {
	Ctx			*SContext
	Cost		int
	Zero		SVertex							// coordinates of zero square
	Puzzle		[][]Tnumber						// puzzle grid
	Path		[]Tnumber
}

func (node SNode) String() string {
	res := "---------------Snode---------------\n"
	res += "Cost : " + fmt.Sprint(node.Cost) + "\n"
	// res += "Size : " + fmt.Sprint(node.Ctx.NSize) + "\n"
	// res += "Final : " + fmt.Sprint(node.Ctx.Final) + "\n"
	res += "Puzzle: {\n"
	formatPuzzle := "%02d "
	for _, tab := range node.Puzzle {
		res += "\t"
		for _, nb := range tab {
			res += fmt.Sprintf(formatPuzzle, nb)
		}
		res += "\n"
	}
	res += "}\nZero : " + fmt.Sprint(node.Zero) + "\n"
	res += "Path: " + fmt.Sprint(node.Path)
	return res
}
