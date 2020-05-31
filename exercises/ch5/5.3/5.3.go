//Write a function to print the contents of all text nodes in an HTML document
//tree. Do not descend into <script> or <style> elements, since their contents are not visible
//in a web browser.
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

	for _, text := range printText(nil, doc) {
		fmt.Println(text)
	}
}

func printText(text []string, n *html.Node) []string {
	if n.Type == html.TextNode {
		text = append(text, n.Data)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		printText(text, c)
	}
	return text
}
