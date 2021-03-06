package structures

import (
	"fmt"
)

type Tnumber uint8
type Tpuzzle [][]Tnumber

type SContext struct {
	Heuristic		func([][]Tnumber, []SVertex)int	// pointer to the heuristic function choosed by user
	Final			[]SVertex						// slice filled with coordinates of completed puzzle
	NSize			Tnumber							// size of the puzzle (NSize * NSize)
	Puzzle			Tpuzzle 						// puzzle grid
	Zero			SVertex							// coordinates of zero square
}

func (puzzle Tpuzzle) String() string {
	NSize := len(puzzle)
	res := "Puzzle: {\n"
	formatPuzzle := "%0" + fmt.Sprint(len(fmt.Sprint(NSize * NSize))) + "d "
	for _, tab := range puzzle {
		res += "\t"
		for _, nb := range tab {
			res += fmt.Sprintf(formatPuzzle, nb)
		}
		res += "\n"
	}
	return res + "}"
}

func (ctx SContext) String() string {
	res := "---------------SContext---------------\n"
	res += "Zero: " + fmt.Sprint(ctx.Zero) + "\n"
	res += "NSize: " + fmt.Sprint(ctx.NSize) + "\n"
	res += fmt.Sprint(ctx.Puzzle)
	res += "\nFinal: " + fmt.Sprint(ctx.Final)
	res += "\n--------------------------------------"
	return res
}