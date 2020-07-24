package solver

import (
	s "n-puzzle/structures"
	"fmt"
)

const (
	NORTH = 1 << iota
	SOUTH = 1 << iota
	WEST = 1 << iota
	EAST = 1 << iota
)

func whereToGo(puzzle [][]uint16, X, Y, NSize, last uint16) (uint8, uint8) {
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

func error(path []uint16) bool {
	for _, nb := range path {
		if nb == 0 {
			return true
		}
	}
	return false
}

func Algo(node *s.SNode) {
	Start:
	// fmt.Println(node)
	// if node.Cost > 5 {
	// 	panic("exit")
	// }
	if node.Cost > 100 {
		panic(fmt.Sprintln(node))
	}
	heuristic := node.Ctx.Heuristic(node.Puzzle, node.Ctx.Final)
	// fmt.Println(heuristic, "+", node.Cost, node.Path)
	if heuristic == 0 {
		node.Ctx.ResultChan <- s.SResult{Res: PathCopy(node.Path)}
	}
	FromMasterChan := make(chan int8, 1)
	// fmt.Println("send request")
	node.Ctx.RequestChan <- s.SRequest{Score: heuristic + node.Cost, Respond: FromMasterChan}
	response := <-FromMasterChan
	// fmt.Println("request response", response)
	if response == -1 {
		return
	}
	node.Cost++
	nswe, pathNb := whereToGo(node.Puzzle, node.Zero.X, node.Zero.Y, node.Ctx.NSize, node.Path[len(node.Path) - 1])
	for pathNb > 1 {
		newNode := CopyNode(node)
		// fmt.Printf("%p\n", &newNode)
		pathNb -= 1
		Y, X := newNode.Zero.Y, newNode.Zero.X
		if nswe & NORTH != 0 {
			nswe -= NORTH
			newNode.Zero.Y--
		} else if nswe & SOUTH != 0 {
			nswe -= SOUTH
			newNode.Zero.Y++
		} else if nswe & WEST != 0 {
			nswe -= WEST
			newNode.Zero.X--
		} else if nswe & EAST != 0 {
			nswe -= EAST
			newNode.Zero.X++
		}
		newNode.Path = append(node.Path, newNode.Puzzle[newNode.Zero.Y][newNode.Zero.X])
		newNode.Puzzle[Y][X], newNode.Puzzle[newNode.Zero.Y][newNode.Zero.X] = newNode.Puzzle[newNode.Zero.Y][newNode.Zero.X], newNode.Puzzle[Y][X]
		go Algo(newNode)
	}
	Y, X := node.Zero.Y, node.Zero.X
	if nswe & NORTH != 0 {
		node.Zero.Y--
	} else if nswe & SOUTH != 0 {
		node.Zero.Y++
	} else if nswe & WEST != 0 {
		node.Zero.X--
	} else if nswe & EAST != 0 {
		node.Zero.X++
	}
	// fmt.Printf("X1: %v, Y1: %v | X2: %v, Y2: %v\n", X, Y, node.Zero.X, node.Zero.Y)
	node.Path = append(node.Path, node.Puzzle[node.Zero.Y][node.Zero.X])
	node.Puzzle[Y][X], node.Puzzle[node.Zero.Y][node.Zero.X] = node.Puzzle[node.Zero.Y][node.Zero.X], node.Puzzle[Y][X]
	goto Start
	return
}
