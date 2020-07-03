package main
import (
	// "fmt"
	"os"
	"bufio"
	"strconv"
	"strings"
	n "n-puzzle/nstruct"
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

func getNSize(scanner *bufio.Scanner) (int, string) {
    scanner.Scan()
    if err := scanner.Err(); err != nil {
        return 0, ReadError
    }
	line := scanner.Text()
	line = epur(line)
	tab := strings.Split(line, " ")
	// fmt.Println(tab, len(tab))
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
	// res := make([][]int, NSize)
	// for scanner.Scan() {
	// 	line := epur(scanner.Text())
	// 	if len(line) < 1 {
	// 		continue
	// 	}
	// 	tab := strings.Split(line, " ")
	// }
	return nil, ""
} 

func ParseFile(filename string) (*n.SContext, string) {
	ctx := &n.SContext{FileName: filename}
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
	}
	ctx.Puzzle, err2 = getPuzzle(scanner, ctx.NSize)
	if err2 != "" {
		return ctx, err2
	}
    if err = scanner.Err(); err != nil {
        return ctx, ReadError
    }
	return ctx, ""
}