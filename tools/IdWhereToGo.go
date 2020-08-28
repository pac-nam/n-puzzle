package tools

import (
	s "n-puzzle/structures"
	"sort"
	// "fmt"
)

func IdWhereToGo(ctx *s.SContext, m map[string]interface{}) (s.Tsuccessors) {
	Y, X := ctx.Zero.Y, ctx.Zero.X
	var list s.Tsuccessors
	if Y != 0 {
		ctx.Puzzle[Y][X], ctx.Puzzle[Y-1][X] = ctx.Puzzle[Y-1][X], ctx.Puzzle[Y][X]
		puzzleString := PuzzleToString(ctx.Puzzle)
		_, exist := m[puzzleString]
		if !exist {
			list = append(list, s.SSuccessor{
				Heuristic: ctx.Heuristic(ctx.Puzzle, ctx.Final),
				NSWE: NORTH,
				PuzzleString: puzzleString,
				Move: ctx.Puzzle[Y][X],
			})
		}
		ctx.Puzzle[Y][X], ctx.Puzzle[Y-1][X] = ctx.Puzzle[Y-1][X], ctx.Puzzle[Y][X]
	}
	if Y != ctx.NSize - 1 {
		ctx.Puzzle[Y][X], ctx.Puzzle[Y+1][X] = ctx.Puzzle[Y+1][X], ctx.Puzzle[Y][X]
		puzzleString := PuzzleToString(ctx.Puzzle)
		_, exist := m[puzzleString]
		if !exist {
			list = append(list, s.SSuccessor {
				Heuristic: ctx.Heuristic(ctx.Puzzle, ctx.Final),
				NSWE: SOUTH,
				PuzzleString: puzzleString,
				Move: ctx.Puzzle[Y][X],
			})
		}
		ctx.Puzzle[Y][X], ctx.Puzzle[Y+1][X] = ctx.Puzzle[Y+1][X], ctx.Puzzle[Y][X]
	}
	if X != 0 {
		ctx.Puzzle[Y][X], ctx.Puzzle[Y][X-1] = ctx.Puzzle[Y][X-1], ctx.Puzzle[Y][X]
		puzzleString := PuzzleToString(ctx.Puzzle)
		_, exist := m[puzzleString]
		if !exist {
			list = append(list, s.SSuccessor{
				Heuristic: ctx.Heuristic(ctx.Puzzle, ctx.Final),
				NSWE: WEST,
				PuzzleString: puzzleString,
				Move: ctx.Puzzle[Y][X],
			})
		}
		ctx.Puzzle[Y][X], ctx.Puzzle[Y][X-1] = ctx.Puzzle[Y][X-1], ctx.Puzzle[Y][X]
	}
	if X != ctx.NSize - 1 {
		ctx.Puzzle[Y][X], ctx.Puzzle[Y][X+1] = ctx.Puzzle[Y][X+1], ctx.Puzzle[Y][X]
		puzzleString := PuzzleToString(ctx.Puzzle)
		_, exist := m[puzzleString]
		if !exist {
			list = append(list, s.SSuccessor{
				Heuristic: ctx.Heuristic(ctx.Puzzle, ctx.Final),
				NSWE: EAST,
				PuzzleString: puzzleString,
				Move: ctx.Puzzle[Y][X],
			})
		}
		ctx.Puzzle[Y][X], ctx.Puzzle[Y][X+1] = ctx.Puzzle[Y][X+1], ctx.Puzzle[Y][X]
	}
	sort.Sort(list)
	return list
}