// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 136.

// The toposort program prints the nodes of a DAG in topological order.
package main

import (
	"fmt"
)

// !+table
// prereqs maps computer science courses to their prerequisites.
var prereqs = map[string][]string{
	// "algorithms":     {"data structures"},
	// "calculus":       {"linear algebra"},
	// "linear algebra": {"calculus"},

	// "compilers": {
	// 	"data structures",
	// 	"formal languages",
	// 	"computer organization",
	// },

	// "data structures":       {"discrete math"},
	// "databases":             {"data structures"},
	// "discrete math":         {"intro to programming"},
	// "formal languages":      {"discrete math"},
	// "networks":              {"operating systems"},
	// "operating systems":     {"data structures", "computer organization"},
	// "programming languages": {"data structures", "computer organization"},
	"b": {"a"},
	"c": {"a", "b"},
	"d": {"b"},
	"e": {"a", "d"},
	"f": {"g"},
}

//!-table

// !+main
func main() {
	res, err := topoSort(prereqs)
	if err != nil {
		fmt.Printf("%v", err)
		return
	}
	for i, course := range res {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
}

func topoSort(m map[string][]string) ([]string, error) {
	var order []string
	seen := make(map[string]string)
	var visitAll func(items map[string]struct{}) error

	visitAll = func(items map[string]struct{}) error {
		for item := range items {
			if _, has := seen[item]; !has {
				seen[item] = "visiting"
				next := map[string]struct{}{}
				for _, dep := range m[item] {
					next[dep] = struct{}{}
				}
				err := visitAll(next)
				if err != nil {
					return err
				}

				seen[item] = "done"
				order = append(order, item)
			} else if seen[item] == "visiting" {
				return fmt.Errorf("has circle in dependencies\n")
			}
		}
		return nil
	}
	keys := map[string]struct{}{}
	for key := range m {
		keys[key] = struct{}{}
	}

	err := visitAll(keys)
	if err != nil {
		return nil, err
	}

	return order, nil
}

//!-main
