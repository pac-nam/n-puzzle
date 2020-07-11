package parse

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
	m "n-puzzle/messages"
	s "n-puzzle/structures"
)

func getNSize(scanner *bufio.Scanner) (int, string) {
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
	return nb, ""
}

func getPuzzle(scanner *bufio.Scanner, ctx *s.SContext) string {
	ctx.Puzzle = make([][]int, ctx.NSize)
	index := -1
	for scanner.Scan() {
		line := epur(scanner.Text())
		if len(line) < 1 {
			continue
		}
		tab := strings.Split(line, " ")
		index += 1
		if len(tab) != ctx.NSize {
			return "Expecting " + fmt.Sprint(ctx.NSize) + " numbers on each line"
		} else if index >= ctx.NSize {
			return "Expecting " + fmt.Sprint(ctx.NSize) + " lines"
		}
		ctx.Puzzle[index] = make([]int, ctx.NSize)
		for j, str := range tab {
			nb, err := strconv.Atoi(str)
			if err != nil {
				return m.AtoiError
			}
			ctx.Puzzle[index][j] = nb
			if nb == 0 {
				ctx.Zero = s.SVertex{Y: index, X: j}
			}
		}
	}
	if index+1 != ctx.NSize {
		return "Expecting " + fmt.Sprint(ctx.NSize) + " lines"
	}
	return ""
}