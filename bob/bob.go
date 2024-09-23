package bob

import (
	"strings"
	"unicode"
)

type Validator struct {
	word string
}

func (v *Validator) IsAllUpper() bool {
	return strings.ToUpper(v.word) == v.word && v.HasLatter()
}

func (v *Validator) HasLatter() bool {
	for _, letter := range v.word {
		if unicode.IsLetter(letter) {
			return true
		}
	}
	return false
}

func (v *Validator) IsAQuestion() bool {
	return strings.HasSuffix(v.word, "?")
}

func (v *Validator) IsEmpty() bool {
	return v.word == ""
}

func NewValidator(word string) *Validator {
	return &Validator{
		word: strings.TrimSpace(word),
	}
}

// Hey is a function that returns a response for Bob
func Hey(remark string) string {
	v := NewValidator(remark)

	if v.IsEmpty() {
		return "Fine. Be that way!"
	} else if v.IsAQuestion() && v.IsAllUpper() {
		return "Calm down, I know what I'm doing!"
	} else if v.IsAQuestion() {
		return "Sure."
	} else if v.IsAllUpper() {
		return "Whoa, chill out!"
	}
	return "Whatever."
}
