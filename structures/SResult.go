package structures

import (
	"fmt"
)

type SResult struct {
	TimeComplexity	uint
	SizeComplexity	uint
	Sequence		[]Tnumber
}

func (res SResult) String() string {
	return fmt.Sprintf("Time Complexity: %T\nSize Complexity: %T\nSequence length: %T\n Sequence: %T\n",
	res.TimeComplexity,
	res.SizeComplexity,
	len(res.Sequence),
	res.Sequence)
}
