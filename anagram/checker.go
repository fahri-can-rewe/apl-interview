package anagram

import "sort"

// Checker reports whether two words are anagrams under the project’s domain rules.
// Preconditions (must be satisfied by the caller):
//   - Inputs are validated via anagram.ValidateWordPair (letters-only per unicode.IsLetter; equal length).
//   - Inputs are already case-normalized (lower-cased using Unicode case folding) in a consistent Unicode normalization form
type Checker interface {
	AreAnagrams(w1, w2 string) bool
}

type FreqMapChecker struct{}

func (FreqMapChecker) AreAnagrams(w1, w2 string) bool {
	return EqualFreqMaps(CountLetters(w1), CountLetters(w2))
}

type SortChecker struct{}

func (SortChecker) AreAnagrams(w1, w2 string) bool {
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
