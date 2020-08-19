package tools


import s "n-puzzle/structures"

func CopyPuzzle(puzzle [][]s.Tnumber, Size s.Tnumber) s.Tpuzzle {
    newPuzzle := make([][]s.Tnumber, Size)
    for Y, line := range puzzle {
        newPuzzle[Y] = make([]s.Tnumber, Size)
        for X, nb := range line {
            newPuzzle[Y][X] = nb
        }
    }
    return newPuzzle
}