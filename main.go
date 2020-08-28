package main

import (
	"fmt"
	s "n-puzzle/structures"
	p "n-puzzle/parse"
	solve "n-puzzle/solver"
	t "n-puzzle/tools"
)

func main() {
	ctx, opt, err := p.Parser()
	if err != "" {
		fmt.Println("error: " + err)
		return
	}
	res := s.SResult{}
	if *opt.IdAStar == true {
		res = solve.IdAstar(ctx)
	} else {
		res = solve.Astar(ctx)
	}
	fmt.Println(res)
	if *opt.Checker == true || *opt.VChecker == true {
		t.Checker(ctx, res.Sequence, *opt.VChecker)
	}
}
