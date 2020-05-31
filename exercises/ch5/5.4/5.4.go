//Extend the visit func tion so that it ext racts other kinds of lin ks from the document,
//such as images, scripts, and style sheets.
//run with fetch
//./fetch https://golang.org | ./5.4.exe
package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}

	links := make(map[string][]string)
	visit(links, doc)
	for key, val := range links {
		fmt.Printf("%d\t%s\t %v\n", len(val), key, val)
	}

}

// visit appends to links each link found in n and returns the result.
func visit(links map[string][]string, n *html.Node) {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links[n.Data] = append(links[n.Data], a.Val)
			}
		}
	}

	if n.Type == html.ElementNode && (n.Data == "img" || n.Data == "script" || n.Data == "link") {
		for _, a := range n.Attr {
			if a.Key == "src" {
				links[n.Data] = append(links[n.Data], a.Val)
			}
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		visit(links, c)
	}
}
