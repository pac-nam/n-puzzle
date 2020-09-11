package structures

import (
	"fmt"
)

type SOption struct {
	FileName		string						// string passed in first command line argument
	HeuristicName	*string
	Checker			*bool
	VChecker		*bool
	IdAStar			*bool
	Help			*bool
}

func (opt SOption) String() string {
	res := "FileName: " + fmt.Sprint(opt.FileName) + "\n"
	res += "Heuristic: " + fmt.Sprint(*opt.HeuristicName) + "\n"
	res += "Checker: " + fmt.Sprint(*opt.Checker) + "\n"
	res += "Visual Checker: " + fmt.Sprint(*opt.VChecker) + "\n"
	res += "Id A*: " + fmt.Sprint(*opt.IdAStar)
	return res
}