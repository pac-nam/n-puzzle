package solver

import (
	"fmt"
	s "n-puzzle/structures"
	t "n-puzzle/tools"
)

func IdAstar(ctx *s.SContext) []s.Tnumber {
	bound := ctx.Heuristic(ctx.Puzzle, ctx.Final)
	CopyPuzzle := t.CopyPuzzle(ctx.Puzzle, ctx.NSize)
	fmt.Println(ctx)
	m := make(map[string]interface{})
	for {
		path, score := search(ctx, 0, bound, m, s.SSuccessor{
			Heuristic: ctx.Heuristic(ctx.Puzzle, ctx.Final),
			NSWE: 0,
			PuzzleString: t.PuzzleToString(ctx.Puzzle),
		})
		if score == 0 {
			// fmt.Println(path[1:], len(path))
			ctx.Puzzle = CopyPuzzle
			path = path[1:]
			for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
				path[i], path[j] = path[j], path[i]
			}
			return path
		}
		bound = score
	}
	return []s.Tnumber{}
}

func search(ctx *s.SContext, cost, costMax int, m map[string]interface{}, father s.SSuccessor) ([]s.Tnumber, int) {
	// fmt.Println(ctx)
	if father.Heuristic == 0 {
		return make([]s.Tnumber, 1), 0
	} else if cost + father.Heuristic > costMax {
		return nil, cost + father.Heuristic
	}
	m[father.PuzzleString] = nil
	min := 10000000000
	neighborgs := t.IdWhereToGo(ctx, m)
	for _, neighborg := range neighborgs {
		m[neighborg.PuzzleString] = nil
		Y, X := ctx.Zero.Y, ctx.Zero.X
		if neighborg.NSWE == t.NORTH {
			ctx.Zero.Y = Y - 1
		} else if neighborg.NSWE == t.SOUTH {
			ctx.Zero.Y = Y + 1
		} else if neighborg.NSWE == t.WEST {
			ctx.Zero.X = X - 1
		} else if neighborg.NSWE == t.EAST {
			ctx.Zero.X = X + 1
		}
		ctx.Puzzle[ctx.Zero.Y][ctx.Zero.X], ctx.Puzzle[Y][X] = ctx.Puzzle[Y][X], ctx.Puzzle[ctx.Zero.Y][ctx.Zero.X]
		path, score := search(ctx, cost + 1, costMax, m, neighborg)
		if score == 0 {
			return append(path, neighborg.Move), 0
		} else if score < min {
			min = score
		}
		ctx.Puzzle[ctx.Zero.Y][ctx.Zero.X], ctx.Puzzle[Y][X] = ctx.Puzzle[Y][X], ctx.Puzzle[ctx.Zero.Y][ctx.Zero.X]
		ctx.Zero.Y, ctx.Zero.X = Y, X
		delete(m, neighborg.PuzzleString)
	}
	return nil, min
}