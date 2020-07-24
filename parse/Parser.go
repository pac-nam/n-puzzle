package parse

import (
	m "n-puzzle/messages"
	s "n-puzzle/structures"
	"os"
)

func Parser() (*s.SContext, string) {
	ctx := &s.SContext{}
	if len(os.Args) != 3 {
		return ctx, m.Help
	}
	err := getHeuristic(ctx, os.Args[1])
	if err != "" {
		return ctx, err
	}
	return parseFile(ctx, os.Args[2])
}
