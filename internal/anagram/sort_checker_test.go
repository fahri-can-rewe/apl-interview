package anagram

import "testing"

func TestSortChecker_AreAnagrams(t *testing.T) {
	runCheckerTests(t, SortChecker{})
}
