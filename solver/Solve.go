package solver

import (
	s "n-puzzle/structures"
	l "container/list"
	"fmt"
)

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func ManhattanHeuristic(puzzle [][]int, final []s.SVertex) int {
	// fmt.Println("hello i am manhattan")
	score := 0
	for Y, line := range puzzle {
		for X, nb := range line {
			score += abs(final[nb].X - X) + abs(final[nb].Y - Y)
		}
	}
	return score
}

func HeuristicPlacement(puzzle [][]int, final []s.SVertex) int {
	score := 0
	for Y, line := range  puzzle {
		for X, nb := range line {
			if final[nb].X != X || final[nb].Y != Y {
				score ++
			}
		}
	}
	return score
}

func CompleteHeuristic(puzzle [][]int, final []s.SVertex) int {
	for Y, line := range  puzzle {
		for X, nb := range line {
			if final[nb].X != X || final[nb].Y != Y {
				return 1
			}
		}
	}
	return 0
}

func testHeuristic(node s.SNode, lst l.List) l.List {
	heurisXL, heurisXR, heurisYR, heurisYL := -1, -1, -1, -1
	zero := node.Zero
	subLst := l.New()
	lst.PushBack(subLst)
	if zero.X != 0 {
		puzzleXL := NodeCopyPuzzle(node.Puzzle, node.Size)
		tmp := puzzleXL[zero.Y][zero.X]
		puzzleXL[zero.Y][zero.X] = puzzleXL[zero.Y][zero.X - 1]
		puzzleXL[zero.Y][zero.X - 1] = tmp
		heurisXL = node.Heuristic(puzzleXL, node.Final)
		newNode := NewNextNode(puzzleXL, s.SVertex{X : zero.X - 1, Y : zero.Y}, node)
		newNode.ValHeuris = heurisXL
		subLst.PushBack(newNode)
	}
	if zero.X != node.Size - 1 {
		puzzleXR := NodeCopyPuzzle(node.Puzzle, node.Size)
		tmp := puzzleXR[zero.Y][zero.X]
		puzzleXR[zero.Y][zero.X] = puzzleXR[zero.Y][zero.X + 1]
		puzzleXR[zero.Y][zero.X + 1] = tmp
		heurisXR = node.Heuristic(puzzleXR, node.Final)
		newNode := NewNextNode(puzzleXR, s.SVertex{X : zero.X + 1, Y : zero.Y}, node)
		newNode.ValHeuris = heurisXR
		if heurisXL > heurisXR {
			subLst.PushFront(newNode)
		} else {
			subLst.PushBack(newNode)
		}
	}
	if zero.Y != 0 {
		puzzleYL := NodeCopyPuzzle(node.Puzzle, node.Size)
		tmp := puzzleYL[zero.Y][zero.X]
		puzzleYL[zero.Y][zero.X] = puzzleYL[zero.Y - 1][zero.X]
		puzzleYL[zero.Y - 1][zero.X] = tmp
		heurisYL = node.Heuristic(puzzleYL, node.Final)
		newNode := NewNextNode(puzzleYL, s.SVertex{X : zero.X, Y : zero.Y - 1}, node)
		newNode.ValHeuris = heurisYL
		isNil := true
		for e := subLst.Front(); e != nil; e = e.Next() {
			val := e.Value.(s.SNode)
			if (val.ValHeuris >= heurisYL) {
				subLst.InsertBefore(newNode, e)
				isNil = false
				break
			}
		}
		if isNil == true {
			subLst.PushBack(newNode)
		}
	}
	if zero.Y != node.Size - 1 {
		puzzleYR := NodeCopyPuzzle(node.Puzzle, node.Size)
		tmp := puzzleYR[zero.Y][zero.X]
		puzzleYR[zero.Y][zero.X] = puzzleYR[zero.Y + 1][zero.X]
		puzzleYR[zero.Y + 1][zero.X] = tmp
		heurisYR = node.Heuristic(puzzleYR, node.Final)
		newNode := NewNextNode(puzzleYR, s.SVertex{X : zero.X, Y : zero.Y + 1}, node)
		newNode.ValHeuris = heurisYR
		isNil := true
		for e := subLst.Front(); e != nil; e = e.Next() {
			val := e.Value.(s.SNode)
			if (val.ValHeuris >= heurisYR) {
				subLst.InsertBefore(newNode, e)
				isNil = false
				break
			}
		}
		if isNil == true {
			subLst.PushBack(newNode)
		}
	}
	for e := subLst.Front() ; e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
	fmt.Println("ok", heurisXL, heurisXR, heurisYL, heurisYR)
	return lst
}


func Solve(ctx s.SContext) {
	
	lst := l.New()
	// lst.PushFront(node)
	testHeuristic(NewNode(ctx), *lst)
	// for e := lst.Front(); e != nil; e = e.Next() {
	// 	fmt.Println(e.Value)
	// }
	return
}