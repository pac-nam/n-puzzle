package structures

import (
	"fmt"
)

type SContext struct {
	// FileName		string				// string passed in first command line argument
	Heuristic		func(*SContext)int	// pointer to the heuristic function choosed by user
	Puzzle			[][]int 			// puzzle grid
	Final			[]SVertex			// slice filled with coordinates of completed puzzle
	NSize			int					// size of the puzzle (NSize * NSize)
}

func (ctx SContext) String() string {
	res := "---------------SContext---------------\n"
	// res += "FileName: \"" + ctx.FileName + "\"\n"
	// res += fmt.Sprintln("Heuristic:", ctx.Heuristic)
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