package tools

import (
    s "n-puzzle/structures"
	"strconv"
)

func PuzzleToString(puzzle s.Tpuzzle) string {
	res := ""
	for _, line := range puzzle {
		for _, nb := range line {
			res += " " + strconv.Itoa(int(nb))
		}
	}
	return res[1:]
}

func AlreadySeen(puzzle s.Tpuzzle, mSeen map[string]interface{}) bool {
	stringPuzzle := PuzzleToString(puzzle)
	if _, exist := mSeen[stringPuzzle]; exist {
		return true
	}
	mSeen[stringPuzzle] = nil
	return false
}