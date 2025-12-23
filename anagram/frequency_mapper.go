package anagram

import "unicode"

func isAlphabetic(word string) bool {
	for _, val := range word {
		if !unicode.IsLetter(val) {
			return false
		}
	}
	return true
}

func CountLetters(word string) map[rune]int {
	letterFreqCount := make(map[rune]int)
	for _, val := range word {
		letterFreqCount[val]++
	}
	return letterFreqCount
}

func EqualFreqMaps(freqCountFirstWord, freqCountSecondWord map[rune]int) bool {
	for char, count := range freqCountFirstWord {
		if freqCountSecondWord[char] != count {
			return false
		}
	}
	return true
}
