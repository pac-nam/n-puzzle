package solver

import (
	"fmt"
	s "n-puzzle/structures"
	t "n-puzzle/tools"
	h "n-puzzle/heuristic"
)

func goodPlace(stack []s.SImage, newElem s.SImage) []s.SImage {
	length := len(stack) - 1
	if length == -1 {
        return []s.SImage{newElem}
    }
    if stack[0].Score > newElem.Score {
        return append([]s.SImage{newElem}, stack...)
    }
    for i, elem := range stack {
        if i == length {
            break
        }
        if elem.Score > newElem.Score {
			tmp := append(stack[:i], newElem)
			return append(tmp, stack[i+1:]...)
			
        }
    }
    return append(stack, newElem)
}

func whereIsZero(puzzle [][]s.Tnumber) s.SVertex {
	for y, line := range puzzle {
		for x, item := range line {
			if item == 0 {
				return(s.SVertex{Y: s.Tnumber(y), X: s.Tnumber(x)})
			}
		}
	}
	return (s.SVertex{Y: s.Tnumber(0), X: s.Tnumber(0)})
}

func exploreNeighborg(image s.SImage, zero s.SVertex, ctx *s.SContext) []s.SImage {
	puzzle := image.Puzzle
	nswe, pathNb := whereToGo(puzzle, zero.X, zero.Y, ctx.NSize, 0)
	neighborg := make([]s.SImage, pathNb)
	for i := 0; nswe != 0; i++ {
		newP := t.CopyPuzzle(puzzle, ctx.NSize)
		Y, X := zero.Y, zero.X
		if nswe & NORTH != 0 {
			nswe -= NORTH
			newP[Y-1][X], newP[Y][X] = newP[Y][X], newP[Y-1][X]
		} else if nswe & SOUTH != 0 {
			nswe -= SOUTH
			newP[Y+1][X], newP[Y][X] = newP[Y][X], newP[Y+1][X]
		} else if nswe & WEST != 0 {
			nswe -= WEST
			newP[Y][X-1], newP[Y][X] = newP[Y][X], newP[Y][X-1]
		} else if nswe & EAST != 0 {
			nswe -= EAST
			newP[Y][X+1], newP[Y][X] = newP[Y][X], newP[Y][X+1]
			
		}
		newImage := s.SImage {Cost : image.Cost, Puzzle : newP, Score : ctx.Heuristic(newP, ctx.Final) + image.Cost}
		neighborg[i] = newImage
	}
	return neighborg
}

func compare(puzzle1, puzzle2[][]s.Tnumber) bool {
	for y, line := range puzzle1 {
		for x, item := range line {
			if puzzle2[y][x] != item {
				return false
			}
		}
	}
    return true
}

func find(slice []s.SImage, puzzle [][]s.Tnumber) (bool, int) {
    for i, item := range slice {
        if item.Puzzle != nil {
			if (compare(item.Puzzle, puzzle) == true) {
				return true, i
			}
        }
    }
    return false, -1
}

func Astar(ctx *s.SContext) {
	image := s.SImage {
		Cost : 0,
		Puzzle : t.CopyPuzzle(ctx.Puzzle, ctx.NSize),
		Score : 0,
	}
	fmt.Println(image)
	opened := make([]s.SImage, 0)
	opened = append(opened, image)
	closed := make([]s.SImage, 0)
	success := false
	for len(opened) != 0 && success != true {
		u := opened[0]
		heuris := ctx.Heuristic(u.Puzzle, ctx.Final)
		if h.Complete(u.Puzzle, ctx.Final) == 0 {
			success = true
			fmt.Println("Trouvé")
			fmt.Println(u)
		} else {
			opened = opened[1:]
			// fmt.Println(opened)
			closed = append(closed, u)
			zero := whereIsZero(u.Puzzle)
			neighborgs := exploreNeighborg(u, zero, ctx)
			for _, neighborg := range neighborgs {
				inClosed, iInClosed := find(closed, neighborg.Puzzle)
				inOpened, _ := find(opened, neighborg.Puzzle)
				if inOpened == false && inClosed == false {
					neighborg.Cost = u.Cost + 1
					opened = goodPlace(opened, neighborg)
				} else {
					heuris = ctx.Heuristic(neighborg.Puzzle, ctx.Final)
					neighborg.Cost = u.Cost + 1
					if heuris + neighborg.Cost > u.Cost + heuris + 1 {
						neighborg.Cost = u.Cost + 1
						if inClosed == true {
							opened = goodPlace(opened, closed[iInClosed])
							closed = append(closed[:iInClosed], closed[iInClosed+1:]...)
						}
					}
				}
			}
		}
	}
	// fmt.Println(opened, success)
	// fmt.Println(find(opened, ctx.Puzzle))
	// opened = append(opened, ctx.Puzzle)
	// fmt.Println(opened)
	// opened = opened[:1]
	// fmt.Println(opened)
}
