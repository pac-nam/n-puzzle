package structures

import (
	"fmt"
)

type Tnumber uint8
type Tpuzzle [][]Tnumber

type SContext struct {
	// FileName		string						// string passed in first command line argument
	Heuristic		func([][]Tnumber, []SVertex)int	// pointer to the heuristic function choosed by user
	Final			[]SVertex					// slice filled with coordinates of completed puzzle
	NSize			Tnumber							// size of the puzzle (NSize * NSize)
	// ResultChan		chan SResult
	// RequestChan		chan SRequest
	Puzzle			[][]Tnumber 					// puzzle grid
	Zero			SVertex						// coordinates of zero square
}

func (ctx SContext) String() string {
	res := "---------------SContext---------------\n"
	// res += "FileName: \"" + ctx.FileName + "\"\n"
	// res += fmt.Sprintln("Heuristic:", ctx.Heuristic)
	res += "Zero: " + fmt.Sprint(ctx.Zero) + "\n"
	res += "NSize: " + fmt.Sprint(ctx.NSize) + "\n"
	res += "Puzzle: {\n"
	formatPuzzle := "%0" + fmt.Sprint(len(fmt.Sprint(ctx.NSize * ctx.NSize))) + "d "
	for _, tab := range ctx.Puzzle {
		res += "\t"
		for _, nb := range tab {
			res += fmt.Sprintf(formatPuzzle, nb)
		}
		res += "\n"
	}
	res += "}\nFinal: " + fmt.Sprint(ctx.Final)
	res += "\n--------------------------------------"
	return res
}