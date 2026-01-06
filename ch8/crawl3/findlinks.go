// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 243.

// Crawl3 crawls web links starting with the command-line arguments.
//
// This version uses bounded parallelism.
// For simplicity, it does not address the termination problem.
package main

import (
	"fmt"
	"log"
	"os"

	"gopl.io/ch5/links"
)

type urls struct {
	list  []string
	depth int
}

func crawl(link string, depth int) urls {
	fmt.Println(link, "  depth:", depth)
	list, err := links.Extract(link)
	if err != nil {
		log.Print(err)
	}
	return urls{list, depth + 1}
}

// !+
func main() {
	worklist := make(chan urls) // lists of URLs, may have duplicates
	unseenLinks := make(chan struct {
		name  string
		depth int
	}) // de-duplicated URLs

	// Add command-line arguments to worklist.
	go func() { worklist <- urls{os.Args[1:], 0} }()

	// Create 20 crawler goroutines to fetch each unseen link.
	for i := 0; i < 20; i++ {
		go func() {
			for link := range unseenLinks {
				foundLinks := crawl(link.name, link.depth)
				go func() { worklist <- foundLinks }()
			}
		}()
	}

	// The main goroutine de-duplicates worklist items
	// and sends the unseen ones to the crawlers.
	seen := make(map[string]bool)
	for list := range worklist {
		depth := list.depth
		if depth > 3 {
			continue
		}
		for _, link := range list.list {
			if !seen[link] {
				seen[link] = true
				unseenLinks <- struct {
					name  string
					depth int
				}{link, depth}
			}
		}
	}
}

//!-
