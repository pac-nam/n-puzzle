package parse

import (
	m "n-puzzle/messages"
	s "n-puzzle/structures"
	"flag"
	"fmt"
	"os"
)

func Option() (*s.SOption, string) {
	opt := s.SOption{}
	opt.HeuristicName = flag.String("H", "manhattan", m.Heuristic)
	opt.Checker = flag.Bool("c", false, "Activate simple checker")
	opt.VChecker = flag.Bool("C", false, "Activate visual checker")
	opt.IdAStar = flag.Bool("i", false, "Activate ID A*")
	opt.Help = flag.Bool("h", false, "print this help")
	flag.Parse()
	if *opt.Help == true || len(flag.Args()) != 1 {
		fmt.Print(m.Help)
		os.Exit(0)
	}
	opt.FileName = flag.Args()[0]
	return &opt, ""
}

func Parser() (*s.SContext, *s.SOption, string) {
	ctx := &s.SContext{}
	opt, err := Option()
	if err != "" {
		return ctx, opt, err
	}
	err = getHeuristic(ctx, *opt.HeuristicName)
	if err != "" {
		return ctx, opt, err
	}
	ctx, err = parseFile(ctx, opt.FileName)
	return ctx, opt, err
}
