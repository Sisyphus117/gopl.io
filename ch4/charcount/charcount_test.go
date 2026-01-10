// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package main

import (
	"fmt"
	"strings"
	"testing"
	"unicode/utf8"
)

func TestCharcount(t *testing.T) {
	tests := []struct {
		val  string
		want *CharCountResult
	}{{
		val: "hello",
		want: &CharCountResult{
			counts: map[rune]int{
				'h': 1,
				'e': 1,
				'l': 2,
				'o': 1,
			},
			utflen:  [utf8.UTFMax + 1]int{1: 5},
			invalid: 0,
		},
	}}
	for _, test := range tests {
		res, err := CharCount(strings.NewReader(test.val))
		if err != nil {
			t.Errorf("%+v", err)
			continue
		}
		for key, val := range test.want.counts {
			if res.counts[key] != val {
				t.Errorf("charcount:%q\twant %d, get %d\n", key, val, res.counts[key])
			}
		}
		for i, n := range res.utflen {
			if n != test.want.utflen[i] {
				t.Errorf("len:%d\twant %d, get %d\n", i, test.want.utflen[i], n)
			}
		}
		if res.invalid != test.want.invalid {
			fmt.Printf("want %d invalid UTF-8 characters, get %d\n", test.want.invalid, res.invalid)
		}
	}
}
