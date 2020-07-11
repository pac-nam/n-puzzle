package structures

import (
    "fmt"
)

type SNode struct {
    ValHeuris   int                         // value of current heuristic
    Zero        SVertex                     // coordinates of zero square 
    Size        int                         // size of the puzzle (NSize * NSize)
    Puzzle      [][]int                     // puzzle grid
    Heuristic   func([][]int, []SVertex)int // pointer to the heuristic function choosed by user
    Final       []SVertex                   // slice filled with coordinates of completed puzzle
}

func (node SNode) String() string {
    res := "---------------Snode---------------\n"
    res += "Heuristique : " + fmt.Sprint(node.ValHeuris) + "\n"
    res += "Size : " + fmt.Sprint(node.Size) + "\n"
    res += "Final : " + fmt.Sprint(node.Final) + "\n"
    res += "Puzzle: {\n"
	formatPuzzle := "%03d "
	for _, tab := range node.Puzzle {
		res += "\t"
		for _, nb := range tab {
			res += fmt.Sprintf(formatPuzzle, nb)
		}
		res += "\n"
    }
    res += "}\nZero : " + fmt.Sprint(node.Zero) + "\n"
    return res
}