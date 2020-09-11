package parse

import (
	"bufio"
	"fmt"
	m "n-puzzle/messages"
	s "n-puzzle/structures"
	"strconv"
	"strings"
	"n-puzzle/heuristic"
)

func getHeuristic(ctx *s.SContext, arg string) string {
	arg = strings.ToLower(arg)
	if arg == "manhattan" || arg == "m" || arg == "1" {
		ctx.Heuristic = heuristic.Manhattan
	} else if arg == "haming" || arg == "h" || arg == "2" {
		ctx.Heuristic = heuristic.Hamming
	} else if arg == "euclidian" || arg == "e" || arg == "3" {
		ctx.Heuristic = heuristic.Euclidian
	} else if arg == "complete" || arg == "c" || arg == "4" {
		ctx.Heuristic = heuristic.Complete
	} else {
		return m.InvalidHeuristic
	}
	return ""
}

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
	} else if nb < 0 {
		return 0, "negative number"
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
			} else if nb < 0 {
				return "negative number"
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
