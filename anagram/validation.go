package anagram

import (
	"errors"
	"fmt"
)

func ValidateWordPair(word1, word2 string) error {
	if !isAlphabetic(word1) || !isAlphabetic(word2) {
		return errors.New("words must contain only letters")
	}
	if len([]rune(word1)) != len([]rune(word2)) {
		return fmt.Errorf("words do not match in length, word1: %q & word2: %q", word1, word2)
	}
	return nil
}
