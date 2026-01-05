package ch5

import "strings"

func expand(s string, f func(string) string) string {
	newSub := f("foo")
	for {
		idx := strings.Index(s, "foo")
		if idx == -1 {
			break
		}
		s = s[:idx] + newSub + s[idx+2:]
	}
	return s
}
