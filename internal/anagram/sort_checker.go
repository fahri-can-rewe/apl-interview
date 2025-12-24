package anagram

import "sort"

type SortChecker struct{}

func (SortChecker) AreAnagrams(w1, w2 string) bool {
	err := validateWordPair(w1, w2)
	if err != nil {
		return false
	}
	r1 := []rune(w1)
	r2 := []rune(w2)
	sort.Slice(r1, func(i, j int) bool { return r1[i] < r1[j] })
	sort.Slice(r2, func(i, j int) bool { return r2[i] < r2[j] })
	for i := range r1 {
		if r1[i] != r2[i] {
			return false
		}
	}
	return true
}
