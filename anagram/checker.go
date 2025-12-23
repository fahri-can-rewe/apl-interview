package anagram

import (
	"errors"
	"fmt"
	"unicode"
)

// Checker reports whether two words are anagrams under the project’s domain rules.
// Preconditions (must be satisfied by the caller):
//   - Inputs are validated via letters-only per unicode.IsLetter
//   - Inputs must be equal length and non-empty
type Checker interface {
	AreAnagrams(w1, w2 string) bool
}

func validateWordPair(word1, word2 string) error {
	if word1 == "" || word2 == "" {
		return errors.New("words must be non-empty")
	}
	if !isAlphabetic(word1) || !isAlphabetic(word2) {
		return errors.New("words must contain only letters")
	}
	if len([]rune(word1)) != len([]rune(word2)) {
		return fmt.Errorf("words do not match in length, word1: %q & word2: %q", word1, word2)
	}
	return nil
}

func isAlphabetic(word string) bool {
	for _, val := range word {
		if !unicode.IsLetter(val) {
			return false
		}
	}
	return true
}
