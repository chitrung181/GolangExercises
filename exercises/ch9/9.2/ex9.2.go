//create lookup table using sync.One
package main

import (
	"fmt"
	"os"
	"strconv"
	"sync"
)

var loadTableOne sync.Once
var pc [256]byte

func initTable() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func main() {
	for _, arg := range os.Args[1:] {
		val, err := strconv.ParseUint(arg, 10, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Convert: %s\n", err)
			os.Exit(1)
		}
		fmt.Printf("Number:%d\tPopCount: %d\n", val, popCount(val))
	}
}

// PopCount returns the population count (number of set bits) of x.
func popCount(x uint64) int {
	loadTableOne.Do(initTable)
	sum := 0
	for i := 0; i < 8; i++ {
		sum += int(pc[byte(x>>(i*8))])
	}
	return sum
}
