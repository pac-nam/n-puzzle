package nstruct

import "fmt"

type SPacket struct {
	OpenedStateNumber		uint64	// Total number of states ever selected in the "opened" set (complexity in time)
	MemoryStateNumber		uint64	// Maximum number of states ever represented in memory at the same time during the search (complexity in size)
	MoveNumber				uint64	// Number of moves required to transition from the initial state to the final state, according to the search
	Sequence				[]int	// The ordered sequence of states that make up the solution, according to the search
	Solvable				bool	// The puzzle may be unsolvable
}

func (p SPacket) String() string {
	res := "----------------SPacket----------------\n"
	res += "OpenedStateNumber: " + fmt.Sprint(p.OpenedStateNumber) + "\n"
	res += "MemoryStateNumber: " + fmt.Sprint(p.MemoryStateNumber) + "\n"
	res += "MoveNumber: " + fmt.Sprint(p.MoveNumber) + "\n"
	res += "Sequence: " + fmt.Sprint(p.Sequence) + "\n"
	res += "Solvable: " + fmt.Sprint(p.Solvable) + "\n"
	res += "--------------------------------------"
	return res
}