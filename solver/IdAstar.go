package solver

import (
	s "n-puzzle/structures"
	t "n-puzzle/tools"
)

func IdAstar(ctx *s.SContext) s.SResult {
	bound := ctx.Heuristic(ctx.Puzzle, ctx.Final)
	CopyPuzzle := t.CopyPuzzle(ctx.Puzzle, ctx.NSize)
	res := s.SResult{TimeComplexity: 0, SizeComplexityMax: 1}
	m := make(map[string]interface{})
	for {
		path, score := search(ctx, 0, bound, m, &res,
			s.SSuccessor {
			Heuristic: ctx.Heuristic(ctx.Puzzle, ctx.Final),
			NSWE: 0,
			PuzzleString: t.PuzzleToString(ctx.Puzzle),
		})
		if score == 0 {
			ctx.Puzzle = CopyPuzzle
			res.Sequence = path[1:]
			for i, j := 0, len(res.Sequence)-1; i < j; i, j = i+1, j-1 {
				res.Sequence[i], res.Sequence[j] = res.Sequence[j], res.Sequence[i]
			}
			return res
		}
		bound = score
	}
	return res
}

func search(ctx *s.SContext, cost, costMax int, m map[string]interface{}, res *s.SResult, father s.SSuccessor) ([]s.Tnumber, int) {
	res.TimeComplexity++
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
		path, score := search(ctx, cost + 1, costMax, m, res, neighborg)
		if uint(len(m)) >= res.SizeComplexityMax {
			res.SizeComplexityMax = uint(len(m)) + 1
		}
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