// add ',' to a number
package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	for _, val := range os.Args[1:] {
		fmt.Println(comma(val))
	}
}

func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}

	var buf bytes.Buffer
	count := 0
	for i := len(s) - 1; i >= 0; i-- {
		buf.WriteByte(s[i])
		count++
		if count == 3 {
			buf.WriteString(",")
			count = 0
		}
	}
	return reverse(buf.String())
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
