package solver

import (
	// "fmt"
	s "n-puzzle/structures"
	t "n-puzzle/tools"
	// h "n-puzzle/heuristic"
)

func exploreNorth(image s.SImage, zero s.SVertex, ctx *s.SContext, result chan s.SImage) {
	newP := t.CopyPuzzle(image.Puzzle, ctx.NSize)
	Y, X := zero.Y, zero.X
	newP[Y-1][X], newP[Y][X] = newP[Y][X], newP[Y-1][X]
	result <- s.SImage {Cost : image.Cost + 1, Puzzle : newP, Score : ctx.Heuristic(newP, ctx.Final) + image.Cost}
}

func exploreSouth(image s.SImage, zero s.SVertex, ctx *s.SContext, result chan s.SImage) {
	newP := t.CopyPuzzle(image.Puzzle, ctx.NSize)
	Y, X := zero.Y, zero.X
	newP[Y+1][X], newP[Y][X] = newP[Y][X], newP[Y+1][X]
	result <- s.SImage {Cost : image.Cost + 1, Puzzle : newP, Score : ctx.Heuristic(newP, ctx.Final) + image.Cost}
}

func exploreWest(image s.SImage, zero s.SVertex, ctx *s.SContext, result chan s.SImage) {
	newP := t.CopyPuzzle(image.Puzzle, ctx.NSize)
	Y, X := zero.Y, zero.X
	newP[Y][X-1], newP[Y][X] = newP[Y][X], newP[Y][X-1]
	result <- s.SImage {Cost : image.Cost + 1, Puzzle : newP, Score : ctx.Heuristic(newP, ctx.Final) + image.Cost}
}

func exploreEast(image s.SImage, zero s.SVertex, ctx *s.SContext, result chan s.SImage) {
	newP := t.CopyPuzzle(image.Puzzle, ctx.NSize)
	Y, X := zero.Y, zero.X
	newP[Y][X+1], newP[Y][X] = newP[Y][X], newP[Y][X+1]
	result <- s.SImage {Cost : image.Cost + 1, Puzzle : newP, Score : ctx.Heuristic(newP, ctx.Final) + image.Cost}
}

func exploreNeighborg(image s.SImage, zero s.SVertex, ctx *s.SContext) []s.SImage {
	nswe, pathNb := whereToGo(image.Puzzle, zero.X, zero.Y, ctx.NSize, 0)
	neighborg := make([]s.SImage, pathNb)
	resultChan := make(chan s.SImage, 4)
	for i := 0; nswe != 0; i++ {
		if nswe & NORTH != 0 {
			nswe -= NORTH
			go exploreNorth(image, zero, ctx, resultChan)
		} else if nswe & SOUTH != 0 {
			nswe -= SOUTH
			go exploreSouth(image, zero, ctx, resultChan)
		} else if nswe & WEST != 0 {
			nswe -= WEST
			go exploreWest(image, zero, ctx, resultChan)
		} else if nswe & EAST != 0 {
			nswe -= EAST
			go exploreEast(image, zero, ctx, resultChan)
			
		}
	}
	for pathNb > 0 {
		pathNb--
		neighborg[pathNb] = <-resultChan
	}
	return neighborg
}