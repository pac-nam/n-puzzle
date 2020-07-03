package main

import (
	"fmt"
	"os"
	p "n-puzzle/parse"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println(InvalidArgNumber)
		return
	}
	ctx, err := p.ParseFile(os.Args[1])
	if err != "" {
		fmt.Println("error: " + err)
		return
	}
	// Soleve(ctx)
	fmt.Println(ctx, "Normal end")
}
