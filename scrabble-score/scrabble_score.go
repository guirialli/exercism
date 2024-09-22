package scrabble

import "strings"

type Scrabble map[rune]int

func NewScrabble() Scrabble {
	return Scrabble{
		'a': 1, 'e': 1, 'i': 1, 'o': 1, 'u': 1,
		'l': 1, 'n': 1, 'r': 1, 's': 1, 't': 1,
		'd': 2, 'g': 2,
		'b': 3, 'c': 3, 'm': 3, 'p': 3,
		'f': 4, 'h': 4, 'v': 4, 'w': 4, 'y': 4,
		'k': 5,
		'j': 8, 'x': 8,
		'q': 10, 'z': 10,
	}
}

func Score(word string) int {
	word = strings.ToLower(word)
	sc := NewScrabble()
	score := 0
	for _, r := range word {
		if v, ok := sc[r]; ok {
			score += v
		}
	}
	return score
}
