package wordcount

import (
	"strings"
	"unicode"
)

type Frequency map[string]int

func (f *Frequency) SeparatePharase(word string) []string {
	word = strings.ReplaceAll(word, "!", ".")
	word = strings.ReplaceAll(word, ":", ".")
	word = strings.ReplaceAll(word, "?", ".")
	word = strings.ReplaceAll(word, "\t", ".")
	word = strings.ReplaceAll(word, "\n", ".")
	word = strings.ReplaceAll(word, " ", ".")
	word = strings.ReplaceAll(word, ",", ".")
	word = strings.ReplaceAll(word, "@", ".")
	word = strings.ReplaceAll(word, "&", ".")
	word = strings.ReplaceAll(word, "$", ".")
	word = strings.ReplaceAll(word, "%", ".")
	word = strings.ReplaceAll(word, "^", ".")
	return strings.Split(word, ".")
}

func (f *Frequency) Nomalize(word string) string {
	isInvalidRune := func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsNumber(r)
	}

	if isInvalidRune(rune(word[0])) {
		word = word[1:]
	}
	if isInvalidRune(rune(word[len(word)-1])) {
		word = word[0 : len(word)-1]
	}
	return strings.ToLower(word)
}

func WordCount(phrase string) Frequency {
	f := Frequency{}
	words := f.SeparatePharase(phrase)
	for _, word := range words {
		if word == "" || word == "'" {
			continue
		}
		f[f.Nomalize(word)] += 1
	}
	return f
}
