package anagram

type FreqMapChecker struct{}

func countLetters(word string) map[rune]int {
	letterFreqCount := make(map[rune]int)
	for _, val := range word {
		letterFreqCount[val]++
	}
	return letterFreqCount
}

func equalFreqMaps(freqCountFirstWord, freqCountSecondWord map[rune]int) bool {
	for char, count := range freqCountFirstWord {
		if freqCountSecondWord[char] != count {
			return false
		}
	}
	return true
}

func (FreqMapChecker) AreAnagrams(w1, w2 string) bool {
	err := validateWordPair(w1, w2)
	if err != nil {
		return false
	}
	return equalFreqMaps(countLetters(w1), countLetters(w2))
}
