// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 123.

// Outline prints the outline of an HTML document tree.
package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

// !+
func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "outline: %v\n", err)
		os.Exit(1)
	}
	outline(nil, doc)
	count := make(map[string]int)
	countElement(nil, doc, count)
	for k, v := range count {
		fmt.Println(k, "  ", v)
	}
}

func outline(stack []string, n *html.Node) {
	if n.Type == html.ElementNode {
		stack = append(stack, n.Data) // push tag
		fmt.Println(stack)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		outline(stack, c)
	}
}

func countElement(stack []string, n *html.Node, count map[string]int) {
	if n.Type == html.ElementNode {
		stack = append(stack, n.Data) // push tag
		count[n.Data]++
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		countElement(stack, c, count)
	}
}

//!-
