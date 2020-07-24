package tools

import (
    s "n-puzzle/structures"
	"strconv"
)

func PuzzleToString(puzzle [][]s.Tnumber) string {
	res := ""
	for _, line := range puzzle {
		for _, nb := range line {
			res += " " + strconv.Itoa(int(nb))
		}
	}
	return res[1:]
}