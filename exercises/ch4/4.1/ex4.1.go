package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"math"
	"os"
)

func main() {
	s1 :=[]byte(os.Args[1])
	s2 :=[]byte(os.Args[2])
	fmt.Printf("number of differnt bits: %g\n",math.Abs(float64(shiftBit(s1) -shiftBit(s2))))
	print(os.Arg[1:])
}

func print(arg []string) {
	switch arg[2] {
		case "SHA384":
			c1 := sha512.Sum384([]byte(arg[0]))
			c2 := sha512.Sum384([]byte(arg[1]))
			fmt.Printf("%x\n%x\n", c1, c2)
		case "SHA512":
			c1 := sha512.Sum512([]byte(arg[0]))
			c2 := sha512.Sum512([]byte(arg[1]))
			fmt.Printf("%x\n%x\n", c1, c2)
		default:
			c1 := sha256.Sum256([]byte(arg[0]))
			c2 := sha256.Sum256([]byte(arg[1]))
			fmt.Printf("%x\n%x\n", c1, c2)
	}
}


func shiftBit(x []byte) int {
	sum := 0
	for _, val := range x {
		for i := 0; i < 8; i++ {
			if val&1 == 1 {
				sum++
			}
			val >>= 1
		}
	}
	return sum
}
