//check if 2 string are anagram
package main

import (
	"fmt"
)

func main() {
	fmt.Println(isAnagram("asdfghj", "jhgfdsa"))
}

func isAnagram(str1, str2 string) bool {
	count := make(map[byte]int)
	for i := 0; i < len(str1); i++ {
		count[str1[i]]++
	}
	for j := 0; j < len(str1); j++ {
		count[str2[j]]--
	}
	for _, num := range count {
		if num != 0 {
			return false
		}
	}
	return true
}
