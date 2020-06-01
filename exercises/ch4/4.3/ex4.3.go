//rotate and reverse items in slide
package main

import "fmt"

func main() {
	arr := [...]int{1, 4, 6, 3, 2, 1, 6, 8}
	sld := arr[:]
	fmt.Printf("Rotate: %v\n", rotate(sld, 3))
	reverse(&arr)
	fmt.Printf("Reverse: %v\n", arr)
}

func reverse(s *[8]int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func rotate(s []int, i int) []int {
	i = i % len(s)
	temp := append(s[i:], s[0:i]...)
	copy(s, temp)
	return s
}
