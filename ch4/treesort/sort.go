// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 101.

// Package treesort provides insertion sort using an unbalanced binary tree.
package treesort

import (
	"strconv"
	"strings"
)

// !+
type tree struct {
	value       int
	left, right *tree
}

func (t *tree) String() string {
	if t == nil {
		return "[]"
	}
	str := "["
	q := []*tree{t}
	for len(q) > 0 {
		n := len(q)
		for i := 0; i < n; i++ {
			cur := q[i]
			if cur == nil {
				str += "nil,"
			} else {
				str += strconv.Itoa(cur.value) + ","
				q = append(q, cur.left, cur.right)
			}
		}
		q = q[n:]
	}
	for strings.HasSuffix(str, "nil,") {
		str = str[:len(str)-4]
	}
	str = str[:len(str)-1] + "]"
	return str
}

// Sort sorts values in place.
func Sort(values []int) {
	var root *tree
	for _, v := range values {
		root = add(root, v)
	}
	appendValues(values[:0], root)
}

// appendValues appends the elements of t to values in order
// and returns the resulting slice.
func appendValues(values []int, t *tree) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}

func add(t *tree, value int) *tree {
	if t == nil {
		// Equivalent to return &tree{value: value}.
		t = new(tree)
		t.value = value
		return t
	}
	if value < t.value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}
	return t
}

//!-
