//Modify dup2 to print the names of all files in which each duplicated line occurs.

//Modify dup2 to print the names of all files in which each duplicated line occurs.

package main

import (
	"bufio"
	"fmt"
	"os"
)

type dupFile struct {
	count int
	files []string
}

func main() {
	counts := make(map[string]*dupFile)
	files := os.Args[1:]
	if len(files) == 0 {
		countLine(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}

			countLine(f, counts)
			f.Close()
		}
	}

	for line, val := range counts {
		if val.count > 1 {
			fmt.Printf("%d\t%s\t%v\n", val.count, line, val.files)
		}
	}
}

func countLine(f *os.File, counts map[string]*dupFile) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		key := input.Text()
		if key == "" {
			return
		}
		if _, ok := counts[key]; ok {
			counts[key].count++
			if !isExistInSlide(f.Name(), counts[key].files) {
				counts[key].files = append(counts[key].files, f.Name())
			}
		} else {
			counts[key] = new(dupFile)
			counts[key].count++
			counts[key].files = append(counts[key].files, f.Name())
		}
	}
}

func isExistInSlide(name string, slide []string) bool {
	for _, val := range slide {
		if val == name {
			return true
		}
	}
	return false
}
