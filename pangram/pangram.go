package pangram

import (
	"strings"
)

type Alphabet struct {
}

func NewAlphabet() *Alphabet {
	return &Alphabet{}
}

func (a Alphabet) IncludeAllAlphabet(word string) bool {
	word = strings.ToLower(word)
	for c := 'a'; c <= 'z'; c++ {
		if !strings.ContainsRune(word, c) {
			return false
		}
	}
	return true

}

func IsPangram(input string) bool {
	return NewAlphabet().IncludeAllAlphabet(input)
}
