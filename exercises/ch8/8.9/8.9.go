//Write a version of du that computes and periodically displays separate totals for
//each of the root directories.
//go run ch8/8.9/8.9.go -v C:\Go C:\Go\src
package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
	"time"
)

var vFlag = flag.Bool("v", false, "show verbose preogress")

type response struct {
	root int
	size int64
}

func main() {
	flag.Parse()

	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}

	responses := make(chan response)
	var wg sync.WaitGroup
	for i, root := range roots {
		wg.Add(1)
		go walkDir(root, &wg, i, responses)
	}
	go func() {
		wg.Wait()
		close(responses)
	}()

	var tick <-chan time.Time
	if *vFlag {
		tick = time.Tick(time.Second)
	}

	nfiles := make([]int64, len(roots))
	nbytes := make([]int64, len(roots))
loop:
	for {
		select {
		case sr, ok := <-responses:
			if !ok {
				break loop // sizeResponses was closed
			}
			nfiles[sr.root]++
			nbytes[sr.root] += sr.size
		case <-tick:
			printDiskUsage(roots, nfiles, nbytes)
		}
	}

	printDiskUsage(roots, nfiles, nbytes) // final totals
}

func printDiskUsage(roots []string, nfiles, nbytes []int64) {
	for i, r := range roots {
		fmt.Printf("%10d files  %.3f GB under %s\n", nfiles[i], float64(nbytes[i])/1e9, r)
	}
}

func walkDir(dir string, n *sync.WaitGroup, root int, responses chan<- response) {
	defer n.Done()
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			n.Add(1)
			subdir := filepath.Join(dir, entry.Name())
			go walkDir(subdir, n, root, responses)
		} else {
			responses <- response{root, entry.Size()}
		}
	}
}

// sema is a counting semaphore for limiting concurrency in dirents.
var sema = make(chan struct{}, 20)

// dirents returns the entries of directory dir.
func dirents(dir string) []os.FileInfo {
	sema <- struct{}{}        // acquire token
	defer func() { <-sema }() // release token

	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du: %v\n", err)
		return nil
	}
	return entries
}
