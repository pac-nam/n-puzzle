package parse

import (
	"os"
	"strings"
	s "n-puzzle/structures"
	"n-puzzle/solver"
	m "n-puzzle/messages"
)

func heuristicType(ctx *s.SContext, arg string) string {
	arg = strings.ToLower(arg)
	if arg == "manhattan" || arg == "m" || arg == "1" {
		ctx.Heuristic = solver.ManhattanHeuristic
	} else if arg == "complete" || arg == "c" || arg == "2" {
		ctx.Heuristic = solver.CompleteHeuristic
	} else if arg == "placement" || arg == "p" || arg == "3" {
		ctx.Heuristic = solver.HeuristicPlacement
	} else {
		return m.InvalidHeuristic
	}
	return ""
}

func Parser() (*s.SContext, string) {
	ctx := &s.SContext{}
	if len(os.Args) != 3 {
		return ctx, m.Help
	}
	err := heuristicType(ctx, os.Args[1])
	if err != "" {
		return ctx, err
	}
	return parseFile(ctx, os.Args[2])
}