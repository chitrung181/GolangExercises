//run: go run ex2.2.go val type
package main

import (
	"fmt"
	"os"
	"strconv"

	conv "exercises/ch2/2.2/convert"
)

func main() {
	val, err := strconv.ParseFloat(os.Args[1], 64)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Convert: %s\n", err)
		os.Exit(1)
	}
	valType := ""
	if len(os.Args[1:]) >= 2 {
		valType = os.Args[2]
	}
	conv.Convert(val, valType)
}
