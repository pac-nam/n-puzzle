package structures

import "fmt"

type SVertex struct {
	X		int		// Coordonnee X representing column
	Y		int		// Coordonnee X representing line
}

func (v SVertex) String() string {
	return fmt.Sprintf("{X: %d, Y: %d}", v.X, v.Y)
}