//remove duplicate in slide
package main

import "fmt"

func main() {
	slide := []string{"a", "b", "b", "a", "ss", "ss", "s", "q"}
	slide = removeAdjDup(slide)
	fmt.Println(slide)
}

func removeAdjDup(slide []string) []string {
	count := len(slide) - 1
	i, j := 0, 1
	for count > 0 {
		if slide[i] == slide[j] {
			slide = remove(slide, j)
		} else {
			i++
			j++
		}
		count--
	}
	return slide
}

func remove(slice []string, i int) []string {
	copy(slice[i:], slice[i+1:])
	return slice[:len(slice)-1]
}
