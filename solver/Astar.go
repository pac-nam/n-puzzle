package solver

import (
	"fmt"
	s "n-puzzle/structures"
	t "n-puzzle/tools"
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

func compare(puzzle1, puzzle2 [][]s.Tnumber) bool {
	for y, line := range puzzle1 {
		for x, item := range line {
			if puzzle2[y][x] != item {
				return false
			}
		}
	}
    return true
}

func find(slice []s.SImage, puzzle [][]s.Tnumber) (bool) {
    for _, item := range slice {
        if item.Puzzle != nil {
			if (compare(item.Puzzle, puzzle) == true) {
				return true
			}
        }
    }
    return false
}

func Astar(ctx *s.SContext) {
	image := s.SImage {
		Cost :		0,
		Puzzle :	t.CopyPuzzle(ctx.Puzzle, ctx.NSize),
		Score :		0,
		Zero:		ctx.Zero,
		Move :		0,
		Father :	nil,
	}
	fmt.Println(image)
	opened := make([]s.SImage, 0)
	opened = append(opened, image)
	// closed := make([]s.SImage, 0)
	closed := make(map[string]*s.SClosed)
	for len(opened) != 0 {
		CurrentImage := opened[0]
		opened = opened[1:]
		heuris := ctx.Heuristic(CurrentImage.Puzzle, ctx.Final)
		if heuris == 0 {
			// return CurrentImage
			fmt.Println("Trouvé")
			fmt.Println(CurrentImage)
			currentClosed := CurrentImage.Father
			for currentClosed != nil {
				fmt.Print(" ", currentClosed.Move)
				currentClosed = currentClosed.Father
			}
			return
		} else {
			// fmt.Println(opened)
			father := CoffeeClosed(closed, CurrentImage)
			neighborgs := exploreNeighborg(CurrentImage, father, ctx)
			fmt.Println(len(neighborgs))
			for _, neighborg := range neighborgs {
				_, existInClosed := closed[t.PuzzleToString(neighborg.Puzzle)]
				inOpened := find(opened, neighborg.Puzzle)
				if inOpened == false && existInClosed == false {
					opened = goodPlace(opened, neighborg)
				}
			}
		}
	}
	fmt.Println(opened)
	// fmt.Println(find(opened, ctx.Puzzle))
	// opened = append(opened, ctx.Puzzle)
	// fmt.Println(opened)
	// opened = opened[:1]
	// fmt.Println(opened)
}
