package solver

import (
	// "fmt"
	s "n-puzzle/structures"
	t "n-puzzle/tools"
	// h "n-puzzle/heuristic"
)

func exploreNorth(image s.SImage, ctx *s.SContext, result chan s.SImage, father *s.Sclosed) {
	newP := t.CopyPuzzle(image.Puzzle, ctx.NSize)
	Y, X := image.Zero.Y, image.Zero.X
	newP[Y-1][X], newP[Y][X] = newP[Y][X], newP[Y-1][X]
	result <- s.SImage {
		Cost : image.Cost + 1,
		Puzzle : newP,
		Score : ctx.Heuristic(newP, ctx.Final) + image.Cost,
		Zero : s.SVertex{Y: s.Tnumber(Y-1), X: s.Tnumber(X)},
		Move : newP[Y][X],
		Father : father,
	}
}

func exploreSouth(image s.SImage, ctx *s.SContext, result chan s.SImage, father *s.SClosed) {
	newP := t.CopyPuzzle(image.Puzzle, ctx.NSize)
	Y, X := image.Zero.Y, image.Zero.X
	newP[Y+1][X], newP[Y][X] = newP[Y][X], newP[Y+1][X]
	result <- s.SImage {
		Cost : image.Cost + 1,
		Puzzle : newP,
		Score : ctx.Heuristic(newP, ctx.Final) + image.Cost,
		Zero : s.SVertex{Y: s.Tnumber(Y+1), X: s.Tnumber(X)},
		Move : newP[Y][X],
		Father : father,
	}
}

func exploreWest(image s.SImage, ctx *s.SContext, result chan s.SImage, father *s.SClosed) {
	newP := t.CopyPuzzle(image.Puzzle, ctx.NSize)
	Y, X := image.Zero.Y, image.Zero.X
	newP[Y][X-1], newP[Y][X] = newP[Y][X], newP[Y][X-1]
	result <- s.SImage {
		Cost : image.Cost + 1,
		Puzzle : newP,
		Score : ctx.Heuristic(newP, ctx.Final) + image.Cost,
		Zero : s.SVertex{Y: s.Tnumber(Y), X: s.Tnumber(X-1)},
		Move : newP[Y][X],
		Father : father,
	}
}

func exploreEast(image s.SImage, ctx *s.SContext, result chan s.SImage, father *s.SClosed) {
	newP := t.CopyPuzzle(image.Puzzle, ctx.NSize)
	Y, X := image.Zero.Y, image.Zero.X
	newP[Y][X+1], newP[Y][X] = newP[Y][X], newP[Y][X+1]
	result <- s.SImage {
		Cost : image.Cost + 1,
		Puzzle : newP,
		Score : ctx.Heuristic(newP, ctx.Final) + image.Cost,
		Zero : s.SVertex{Y: s.Tnumber(Y), X: s.Tnumber(X+1)},
		Move : newP[Y][X],
		Father : father,
	}
}

func exploreNeighborg(image s.SImage, father *s.SClosed, ctx *s.SContext) []s.SImage {
	nswe, pathNb := whereToGo(image.Puzzle, image.Zero.X, image.Zero.Y, ctx.NSize, 0)
	neighborg := make([]s.SImage, pathNb)
	resultChan := make(chan s.SImage, 4)
	for i := 0; nswe != 0; i++ {
		if nswe & NORTH != 0 {
			nswe -= NORTH
			go exploreNorth(image, ctx, resultChan, father)
		} else if nswe & SOUTH != 0 {
			nswe -= SOUTH
			go exploreSouth(image, ctx, resultChan, father)
		} else if nswe & WEST != 0 {
			nswe -= WEST
			go exploreWest(image, ctx, resultChan, father)
		} else if nswe & EAST != 0 {
			nswe -= EAST
			go exploreEast(image, ctx, resultChan, father)
			
		}
	}
	for pathNb > 0 {
		pathNb--
		neighborg[pathNb] = <-resultChan
	}
	return neighborg
}