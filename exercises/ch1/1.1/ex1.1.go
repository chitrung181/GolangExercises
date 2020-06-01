package main

import (
	"fmt"
	"os"
)

func main() {
	for i, val := range os.Args {
		fmt.Printf("%d\t%s\n", i, val)
	}
}
