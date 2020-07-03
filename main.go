package main
import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println(InvalidArgNumber)
		return
	}
	ctx, err := ParseFile(os.Args[1])
	if err != "" {
		fmt.Println("error: " + err)
		return
	}
	Soleve(ctx)
	fmt.Println(ctx, "Normal end")
}