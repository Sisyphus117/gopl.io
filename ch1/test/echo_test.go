package main

import (
	"strings"
	"testing"
)

func slow(args []string) {
	s, sep := "", ""
	for _, arg := range args[1:] {
		s += sep + arg
		sep = " "
	}
	_ = s
}

func fast(args []string) {
	_ = strings.Join(args[1:], " ")
}

func BenchmarkEcho(b *testing.B) {
	args := []string{"hello", "world", "this", "is", "a", "test"}
	b.Run("Slow", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			slow(args)
		}
	})
	b.Run("Fast", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			fast(args)
		}
	})
}
