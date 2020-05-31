//Write a function to populate a mapping from element names—p, div, span,
//and so on—to the number of elements with that name in an HTML document tree
//run with fetch
//./fetch https://golang.org | ./5.2.exe

package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Element count: %v\n", err)
		os.Exit(1)
	}
	count := make(map[string]int)
	mapElement(count, doc)
	for key, num := range count {
		fmt.Printf("%s\t%d\n", key, num)
	}
}

func mapElement(count map[string]int, n *html.Node) {
	if n.Type == html.ElementNode {
		count[n.Data]++
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		mapElement(count, c)
	}
}
