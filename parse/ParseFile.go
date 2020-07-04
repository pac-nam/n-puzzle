package parse

import (
	"bufio"
	s "n-puzzle/structures"
	"os"
	"strings"
)

func epur(line string) string {
	line = strings.Split(line, "#")[0]
	toChange := []string{"\r", "\v", "\t", "\f", "\n"}
	for _, replace := range toChange {
		line = strings.Replace(line, replace, " ", -1)
	}
	tab := strings.Split(line, " ")
	res := ""
	for _, str := range tab {
		if len(str) > 0 {
			res += " " + str
		}
	}
	if len(res) > 0 {
		return res[1:]
	}
	return res
}

func ParseFile(filename string) (*s.SContext, string) {
	ctx := &s.SContext{}
	file, err := os.Open(filename)
	if err != nil {
		return ctx, OpenError
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var err2 string
	ctx.NSize, err2 = getNSize(scanner)
	if err2 != "" {
		return ctx, err2
	} else if ctx.NSize < 2 {
		return ctx, TooSmall
	}
	ctx.Puzzle, err2 = getPuzzle(scanner, ctx.NSize)
	if err2 != "" {
		return ctx, err2
	}
	if err = scanner.Err(); err != nil {
		return ctx, ReadError
	}
	return checkPuzzle(ctx)
}
