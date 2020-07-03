package parse

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

func getNSize(scanner *bufio.Scanner) (int, string) {
	scanner.Scan()
	if err := scanner.Err(); err != nil {
		return 0, ReadError
	}
	line := scanner.Text()
	line = epur(line)
	tab := strings.Split(line, " ")
	if tab[0] == "" {
		return getNSize(scanner)
	} else if len(tab) > 1 {
		return 0, InvalidFirstLine
	}
	nb, err := strconv.Atoi(tab[0])
	if err != nil {
		return 0, AtoiError
	}
	return nb, ""
}

func getPuzzle(scanner *bufio.Scanner, NSize int) ([][]int, string) {
	res := make([][]int, NSize)
	index := -1
	for scanner.Scan() {
		line := epur(scanner.Text())
		if len(line) < 1 {
			continue
		}
		tab := strings.Split(line, " ")
		index += 1
		if len(tab) != NSize {
			return res, "Expecting " + fmt.Sprint(NSize) + " numbers on each line"
		} else if index >= NSize {
			return res, "Expecting " + fmt.Sprint(NSize) + " lines"
		}
		res[index] = make([]int, NSize)
		for j, str := range tab {
			nb, err := strconv.Atoi(str)
			if err != nil {
				return res, AtoiError
			}
			res[index][j] = nb
		}
	}
	if index+1 != NSize {
		return res, "Expecting " + fmt.Sprint(NSize) + " lines"
	}
	return res, ""
}