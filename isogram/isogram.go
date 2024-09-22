package isogram

import (
	"strings"
	"unicode"
)

// Valid if has a equal letters in a word
func hasEqualLetter(word string) bool {
	s := strings.ToLower(word)
	for i, c := range s {
		if unicode.IsLetter(c) && strings.ContainsRune(s[i+1:], c) {
			return false
		}
	}
	return true
}

func IsIsogram(word string) bool {
	return hasEqualLetter(word)
}
