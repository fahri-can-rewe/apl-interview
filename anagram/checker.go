package anagram

import "sort"

type Checker interface {
	AreAnagrams(w1, w2 string) bool
}

type FreqMapChecker struct{}

func (FreqMapChecker) AreAnagrams(w1, w2 string) bool {
	if len([]rune(w1)) != len([]rune(w2)) {
		return false
	}
	return AreAnagrams(CountLetters(w1), CountLetters(w2))
}

type SortChecker struct{}

func (SortChecker) AreAnagrams(w1, w2 string) bool {
	r1 := []rune(w1)
	r2 := []rune(w2)
	if len(r1) != len(r2) {
		return false
	}
	sort.Slice(r1, func(i, j int) bool { return r1[i] < r1[j] })
	sort.Slice(r2, func(i, j int) bool { return r2[i] < r2[j] })
	for i := range r1 {
		if r1[i] != r2[i] {
			return false
		}
	}
	return true
}
