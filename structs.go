package main

import (
	"fmt"
)

type SContext struct {
	FileName		string	// String passed in first command line argument
	NSize			int		// Size of the puzzle (NSize * NSize)
	Puzzle			[][]int // puzzle grid
}

func (ctx SContext) String() string {
	res := "---------------SContext---------------\n"
	res += "FileName: \"" + ctx.FileName + "\"\n"
	res += "NSize: " + fmt.Sprint(ctx.NSize) + "\n"
	res += "Puzzle: {\n"
	for _, tab := range ctx.Puzzle {
		res += "\t"
		for _, nb := range tab {
			res += fmt.Sprintf("%4d", nb)
		}
		res += "\n"
	}
	res += "}\n--------------------------------------"
	return res
}