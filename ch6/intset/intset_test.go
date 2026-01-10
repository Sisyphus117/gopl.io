// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package intset

import (
	"fmt"
	"math/rand"
	"testing"
)

func Example_one() {
	//!+main
	var x, y IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	fmt.Println(x.String()) // "{1 9 144}"

	y.Add(9)
	y.Add(42)
	fmt.Println(y.String()) // "{9 42}"

	x.UnionWith(&y)
	fmt.Println(x.String()) // "{1 9 42 144}"

	fmt.Println(x.Has(9), x.Has(123)) // "true false"
	//!-main

	// Output:
	// {1 9 144}
	// {9 42}
	// {1 9 42 144}
	// true false
}

func Example_two() {
	var x IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	x.Add(42)

	//!+note
	fmt.Println(&x)         // "{1 9 42 144}"
	fmt.Println(x.String()) // "{1 9 42 144}"
	fmt.Println(x)          // "{[4398046511618 0 65536]}"
	//!-note

	// Output:
	// {1 9 42 144}
	// {1 9 42 144}
	// {[4398046511618 0 65536]}
}

func TestAdd(t *testing.T) {
	set := &IntSet{}
	tar := make(map[int]bool)

	set.Add(1)
	set.Add(13)
	set.Add(499)
	set.Add(1321)

	tar[1] = true
	tar[13] = true
	tar[499] = true
	tar[1321] = true

	elems := set.Elems()
	for k := range tar {
		word, bit := k/64, uint(k%64)
		if elems[word]&(1<<bit) == 0 {
			t.Errorf("%d didn't add to intset correctly", k)
		}
	}

	for idx, word := range elems {
		if word == 0 {
			continue
		}
		for i := 0; i < 64; i++ {
			if (word&(1<<i)) != 0 && !tar[idx*64+i] {
				t.Errorf("%d added to intset incorrectly", idx*64+i)
			}
		}

	}

}

const seed = 123456
const LIMIT = 1000_000

type set map[int]struct{}

func (s set) SetAdd(x int) {
	s[x] = struct{}{}
}

func generateRandomArray(seed int, size int) []int {
	rng := rand.New(rand.NewSource(int64(seed)))
	arr := make([]int, size)
	for i := 0; i < size; i++ {
		arr[i] = rng.Intn(LIMIT)
	}
	return arr
}
func BenchmarkSetAdd(b *testing.B) {
	set := set(make(map[int]struct{}))
	arr := generateRandomArray(seed, b.N)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		set.SetAdd(arr[i])
	}
}
func BenchmarkAdd(b *testing.B) {
	set := &IntSet{}
	arr := generateRandomArray(seed, b.N)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		set.Add(arr[i])
	}
}
