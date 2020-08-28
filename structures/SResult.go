package structures

import (
	"fmt"
)

type SResult struct {
	TimeComplexity		uint
	SizeComplexityMax	uint
	Sequence		[]Tnumber
}

func (res SResult) String() string {
	return fmt.Sprintf("Time Complexity: %v\nSize Complexity: %v\nSequence length: %v\nSequence: %v",
	res.TimeComplexity,
	res.SizeComplexityMax,
	len(res.Sequence),
	res.Sequence)
}
