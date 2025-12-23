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
	{name: "emoji true", w1: "😀😃", w2: "😃😀", want: false},
	{name: "empty strings", w1: "", w2: "", want: false},
	{name: "with special characters", w1: "listen!", w2: "silent!", want: false},
	{name: "same letters different counts", w1: "aab", w2: "aba", want: true},
	{name: "case sensitive false", w1: "Listen", w2: "Silent", want: false},
	{name: "different letters", w1: "foo", w2: "bar", want: false},
	{name: "another anagram word pair", w1: "domain", w2: "amodin", want: true},
	{name: "different length word pair", w1: "listentome", w2: "silent", want: false},
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

func TestValidateWordPair(t *testing.T) {
	tests := []struct {
		name    string
		args    []string
		wantErr bool
	}{
		{"valid word pair", []string{"listen", "silent"}, false},
		{"non-alphabetic word pair", []string{"listen1", "silent1"}, true},
		{"same length and characters", []string{"silont", "silent"}, false},
		{"different length and characters", []string{"abc", "ab"}, true},
	}

	for _, testCase := range tests {
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			err := validateWordPair(testCase.args[0], testCase.args[1])
			if (err != nil) != testCase.wantErr {
				t.Errorf("validateWordPair(%v) error = %v, wantErr %v", testCase.args, err, testCase.wantErr)
			}
			if err == nil && testCase.wantErr {
				t.Fatalf("buildAPIClient(%q) expected error, got nil", testCase.args)
			}
		})

	}
}
