package solver

import (
	s "n-puzzle/structures"
	t "n-puzzle/tools"
)

func CoffeeClosed(m map[string]*s.SClosed, image s.SImage) *s.SClosed {
	stringPuzzle := t.PuzzleToString(image.Puzzle)
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