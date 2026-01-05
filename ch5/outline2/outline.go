// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 133.

// Outline prints the outline of an HTML document tree.
package main

import (
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

func main() {
	for _, url := range os.Args[1:] {
		outline(url)
	}
}

func outline(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return err
	}

	//!+call
	startElement, endElement := getStartAndEnd()
	forEachNode(doc, startElement, endElement)
	//!-call

	return nil
}

// !+forEachNode
// forEachNode calls the functions pre(x) and post(x) for each node
// x in the tree rooted at n. Both functions are optional.
// pre is called before the children are visited (preorder) and
// post is called after (postorder).
func forEachNode(n *html.Node, pre, post func(n *html.Node) bool) {
	if pre != nil {
		if !pre(n) {
			return
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}

	if post != nil {
		if !post(n) {
			return
		}
	}
}

func ElementByID(id string, n *html.Node, pre, post func(n *html.Node) bool) (node *html.Node) {

	if pre != nil {
		if !pre(n) {
			return
		}
	}

	for _, attr := range n.Attr {
		if attr.Key == "id" && attr.Val == id {
			node = n
			return
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		ElementByID(id, c, pre, post)
	}

	if post != nil {
		if !post(n) {
			return
		}
	}
	return
}

//!-forEachNode

// !+startend

func getStartAndEnd() (func(*html.Node) bool, func(*html.Node) bool) {
	var depth int
	return func(n *html.Node) bool {
			if n.Type == html.ElementNode {
				fmt.Printf("%*s<%s>\n", depth*2, "", n.Data)
				depth++
				return true
			}
			return false
		}, func(n *html.Node) bool {
			if n.Type == html.ElementNode {
				depth--
				fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
				return true
			}
			return false
		}
}

//!-startend
