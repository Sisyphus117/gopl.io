package ch7

import "sort"

//type IsPalindrome string

//func (i IsPalindrome) Len() int { return len(i) }
//func (i IsPalindrome)Swap(x,y int){i[x]}

func IsPalindrome(s sort.Interface) bool {
	n := s.Len()
	for i := 0; i < n/2; i++ {
		if s.Less(i, n-1-i) || s.Less(n-1-i, i) {
			return false
		}
	}
	return true
}
