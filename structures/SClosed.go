package structures

import (
	"fmt"
)

type SClosed struct {
	Cost		int
	Move		Tnumber
	Father		*SClosed
}

func (node SClosed) String() string {
	res := "---------------SClosed---------------\n"
	res += "Cost : " + fmt.Sprint(node.Cost) + "\n"
	return res
}
