package solver

import (
	s "n-puzzle/structures"
)

func CoffeeClosed(m map[string]*s.SClosed, image s.SImage) *s.SClosed {
	stringPuzzle := image.PuzzleString
	elem, exist := m[stringPuzzle]
	if !exist || elem.Cost > image.Cost {
		m[stringPuzzle] = &s.SClosed {
			Cost: image.Cost,
			Move: image.Move,
			Father: image.Father,
		}
	}
	return m[stringPuzzle]
}