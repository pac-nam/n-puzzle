package structures

import (
    "fmt"
)

type SNode struct {
    Final       []SVertex          // slice filled with coordinates of completed puzzle
    Puzzle      [][]int             // puzzle grid
    Heuristic   func([][]int, []SVertex)int  // pointer to the heuristic function choosed by user
}

func (node SNode) String() string {
    res := "---------------Snode---------------\n"
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
    res += "}"
    return res
}