package parse

import (
	"bufio"
	"fmt"
	m "n-puzzle/messages"
	s "n-puzzle/structures"
	"strconv"
	"strings"
)

func getNSize(scanner *bufio.Scanner) (s.Tnumber, string) {
	scanner.Scan()
	if err := scanner.Err(); err != nil {
		return 0, m.ReadError
	}
	line := scanner.Text()
	line = epur(line)
	tab := strings.Split(line, " ")
	if tab[0] == "" {
		return getNSize(scanner)
	} else if len(tab) > 1 {
		return 0, m.InvalidFirstLine
	}
	nb, err := strconv.Atoi(tab[0])
	if err != nil {
		return 0, m.AtoiError
	}
	return s.Tnumber(nb), ""
}

func getPuzzle(scanner *bufio.Scanner, ctx *s.SContext) string {
	ctx.Puzzle = make([][]s.Tnumber, ctx.NSize)
	index := -1
	for scanner.Scan() {
		line := epur(scanner.Text())
		if len(line) < 1 {
			continue
		}
		tab := strings.Split(line, " ")
		index += 1
		if len(tab) != int(ctx.NSize) {
			return "Expecting " + fmt.Sprint(ctx.NSize) + " numbers on each line"
		} else if index >= int(ctx.NSize) {
			return "Expecting " + fmt.Sprint(ctx.NSize) + " lines"
		}
		ctx.Puzzle[index] = make([]s.Tnumber, ctx.NSize)
		for j, str := range tab {
			nb, err := strconv.Atoi(str)
			if err != nil {
				return m.AtoiError
			}
			ctx.Puzzle[index][j] = s.Tnumber(nb)
			if nb == 0 {
				ctx.Zero = s.SVertex{Y: s.Tnumber(index), X: s.Tnumber(j)}
			}
		}
	}
	if index+1 != int(ctx.NSize) {
		return "Expecting " + fmt.Sprint(ctx.NSize) + " lines"
	}
	return ""
}
