package anagram

import "testing"

var areAnagramCases = []struct {
	name string
	w1   string
	w2   string
	want bool
}{
	{name: "simple true", w1: "listen", w2: "silent", want: true},
	{name: "another true", w1: "triangle", w2: "integral", want: true},
	{name: "unicode accents true", w1: "résumé", w2: "ésumér", want: true},
	{name: "cjk true", w1: "夜空", w2: "空夜", want: true},
	{name: "emoji true", w1: "😀😃", w2: "😃😀", want: true},
	{name: "empty strings", w1: "", w2: "", want: true},
	{name: "same letters different counts", w1: "aab", w2: "aba", want: true},
	{name: "case sensitive false", w1: "Listen", w2: "Silent", want: false},
	{name: "different letters", w1: "foo", w2: "bar", want: false},
}

func runCheckerTests(t *testing.T, c Checker) {
	t.Helper()
	for _, tc := range areAnagramCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			if got := c.AreAnagrams(tc.w1, tc.w2); got != tc.want {
				t.Errorf("AreAnagrams(%q, %q) = expected: %v, actual: %v", tc.w1, tc.w2, tc.want, got)
			}
		})
	}
}

func TestFreqMapChecker_AreAnagrams(t *testing.T) {
	runCheckerTests(t, FreqMapChecker{})
}

func TestSortChecker_AreAnagrams(t *testing.T) {
	runCheckerTests(t, SortChecker{})
}
