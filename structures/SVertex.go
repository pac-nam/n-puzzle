package structures

import "fmt"

type SVertex struct {
	Y		Tnumber		// Coordinate X representing line
	X		Tnumber		// Coordinate X representing column
}

func (v SVertex) String() string {
	return fmt.Sprintf("{X: %d, Y: %d}", v.X, v.Y)
}