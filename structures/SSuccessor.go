package structures

// import (
// 	"fmt"
// )

type SSuccessor struct {
	Heuristic		int
	NSWE			int
}

type Tsuccessors []SSuccessor

func (list Tsuccessors)Len() int {
	return len(list)
}

func (list Tsuccessors)Swap(i, j int) {
	list[i], list[j] = list[j], list[i]
}

func (list Tsuccessors)Less(i, j int) bool {
	return list[i].Heuristic < list[j].Heuristic
}
// func (node SSuccessor) String() string {
// 	res := "---------------SImage---------------\n"
// 	res += "Cost : " + fmt.Sprint(node.Cost) + "\n"
// 	// res += "Size : " + fmt.Sprint(node.Ctx.NSize) + "\n"
// 	// res += "Final : " + fmt.Sprint(node.Ctx.Final) + "\n"
// 	res += "Puzzle: {\n"
// 	formatPuzzle := "%02d "
// 	for _, tab := range node.Puzzle {
// 		res += "\t"
// 		for _, nb := range tab {
// 			res += fmt.Sprintf(formatPuzzle, nb)
// 		}
// 		res += "\n"
// 	}
// 	res += "}"
// 	return res
// }