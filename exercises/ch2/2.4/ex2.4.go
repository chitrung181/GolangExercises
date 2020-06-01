//count number of set bits
package main

import (
	"fmt"
	"os"
	"strconv"
)

var pc [256]byte

func init() {
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
		fmt.Printf("Number:%d\nPopCount: %d\nShiftBit: %d", val, popCount(val), shiftBit(val))
	}
}

// PopCount returns the population count (number of set bits) of x.
func popCount(x uint64) int {
	sum := 0
	for i := 0; i < 8; i++ {
		sum += int(pc[byte(x>>(i*8))])
	}
	return sum
}

// shiftBit return number of set bits using shift bit and compare
func shiftBit(x uint64) int {
	sum := 0
	for i := 0; i < 64; i++ {
		if x&1 == 1 {
			sum++
		}
		x >>= 1
	}
	return sum
}
