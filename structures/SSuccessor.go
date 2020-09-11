package structures

// import (
// 	"fmt"
// )

type SSuccessor struct {
	Heuristic		int
	NSWE			int
	PuzzleString	string
	Move			Tnumber
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