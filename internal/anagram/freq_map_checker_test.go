package anagram

import (
	"maps"
	"testing"
)

func TestCountLetters(t *testing.T) {
	tests := []struct {
		word string
		want map[rune]int
	}{
		{word: "hello", want: map[rune]int{'h': 1, 'e': 1, 'l': 2, 'o': 1}},
		{word: "abc", want: map[rune]int{'a': 1, 'b': 1, 'c': 1}},
		{word: "hummus", want: map[rune]int{'h': 1, 'u': 2, 'm': 2, 's': 1}},
		{word: "Batman", want: map[rune]int{'B': 1, 'a': 2, 't': 1, 'm': 1, 'n': 1}},
		{word: "Alpha", want: map[rune]int{'A': 1, 'l': 1, 'p': 1, 'h': 1, 'a': 1}},
		{word: "Celsius", want: map[rune]int{'C': 1, 'e': 1, 'l': 1, 's': 2, 'i': 1, 'u': 1}},
	}

	for _, testCase := range tests {
		testCase := testCase
		t.Run(testCase.word, func(t *testing.T) {
			if got := countLetters(testCase.word); !maps.Equal(got, testCase.want) {
				t.Errorf("CountLetters(%q) = expected: %v, actual: %v", testCase.word, testCase.want, got)
			}
		})
	}
}

func TestEqualFreqMaps(t *testing.T) {
	tests := []struct {
		name   string
		first  map[rune]int
		second map[rune]int
		want   bool
	}{
		{
			name:   "equal maps",
			first:  map[rune]int{'t': 1, 'h': 1, 'o': 1, 'r': 1, 'n': 1, 'e': 1},
			second: map[rune]int{'t': 1, 'h': 1, 'r': 1, 'o': 1, 'n': 1, 'e': 1},
			want:   true,
		},
		{
			name:   "different lengths",
			first:  map[rune]int{'b': 1, 'a': 2, 't': 1, 'm': 1, 'n': 1},
			second: map[rune]int{'b': 1, 'a': 2, 't': 1, 'm': 1},
			want:   false,
		},
		{
			name:   "same length but different keys",
			first:  map[rune]int{'a': 1},
			second: map[rune]int{'b': 1},
			want:   false,
		},
		{
			name:   "same keys different counts",
			first:  map[rune]int{'a': 2, 'b': 1},
			second: map[rune]int{'a': 1, 'b': 1},
			want:   false,
		},
		{
			name:   "both nil maps",
			first:  nil,
			second: nil,
			want:   true,
		},
		{
			name:   "nil vs empty map",
			first:  nil,
			second: map[rune]int{},
			want:   true,
		},
	}

	for _, testCase := range tests {
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			if got := equalFreqMaps(testCase.first, testCase.second); got != testCase.want {
				t.Errorf("EqualFreqMaps(%v, %v) = expected: %v, actual: %v", testCase.first, testCase.second, testCase.want, got)
			}
		})
	}
}

func TestFreqMapChecker_AreAnagrams(t *testing.T) {
	runCheckerTests(t, FreqMapChecker{})
}
