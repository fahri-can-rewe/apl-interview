package anagram

import (
	"maps"
	"testing"
)

func TestIsAlphabetic(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want bool
	}{
		{name: "ascii letters", in: "HelloWorld", want: true},
		{name: "single letter", in: "a", want: true},
		{name: "empty string", in: "", want: true}, // current implementation returns true
		{name: "contains space", in: "hello world", want: false},
		{name: "contains tab", in: "hello\tworld", want: false},
		{name: "contains newline", in: "hello\nworld", want: false},
		{name: "contains digit", in: "abc123", want: false},
		{name: "contains punctuation", in: "abc!", want: false},
		{name: "contains underscore", in: "abc_def", want: false},

		// Unicode letters should be accepted because unicode.IsLetter is used.
		{name: "latin letters with accents", in: "café", want: true},
		{name: "german umlaut", in: "für", want: true},
		{name: "greek letters", in: "Αλφα", want: true},
		{name: "cjk letters", in: "汉字", want: true},

		// Combining mark alone is not a letter (unicode.IsLetter returns false for Mn).
		{name: "combining mark only", in: "\u0301", want: false},
	}

	for _, testCase := range tests {
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			if got := isAlphabetic(testCase.in); got != testCase.want {
				t.Errorf("IsAlphabetic(%q) = expected: %v, actual: %v", testCase.in, testCase.want, got)
			}
		})
	}
}

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
			if got := CountLetters(testCase.word); !maps.Equal(got, testCase.want) {
				t.Errorf("CountLetters(%q) = expected: %v, actual: %v", testCase.word, testCase.want, got)
			}
		})
	}
}

func TestAreAnagrams(t *testing.T) {
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
			if got := EqualFreqMaps(testCase.first, testCase.second); got != testCase.want {
				t.Errorf("EqualFreqMaps(%v, %v) = expected: %v, actual: %v", testCase.first, testCase.second, testCase.want, got)
			}
		})
	}
}
