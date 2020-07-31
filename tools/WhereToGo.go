package tools

import (
	s "n-puzzle/structures"
	// "fmt"
)

const (
	NORTH = 1 << iota
	SOUTH = 1 << iota
	WEST = 1 << iota
	EAST = 1 << iota
)

func WhereToGo(puzzle s.Tpuzzle, X, Y, NSize, last s.Tnumber) (uint8, uint8) {
	var NSWE, pathNb uint8 = 0, 0
	if Y != 0 && puzzle[Y-1][X] != last {
		NSWE |= NORTH
		pathNb++
	}
	if Y != NSize - 1 && puzzle[Y+1][X] != last {
		NSWE |= SOUTH
		pathNb++
	}
	if X != 0 && puzzle[Y][X-1] != last {
		NSWE |= WEST
		pathNb++
	}
	if X != NSize - 1 && puzzle[Y][X+1] != last {
		NSWE |= EAST
		pathNb++
	}
	return NSWE, pathNb
}