package structures

import (
	"fmt"
)

type SImage struct {
	Cost	int
	Score	int
	Puzzle	[][]Tnumber // puzzle grid
	Zero	SVertex
	Move	Tnumber
	// Father	*SClosed
}

func (node SImage) String() string {
	res := "---------------SImage---------------\n"
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
	res += "}"
	return res
}
