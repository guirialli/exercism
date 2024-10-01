package atbash

import (
	"strings"
	"unicode"
)

type Plain []rune

func New() Plain {
	return Plain{
		'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j',
		'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't',
		'u', 'v', 'w', 'x', 'y', 'z',
	}
}

func (p Plain) getElementIndex(r rune) int {
	for i, val := range p {
		if val == r {
			return i
		}
	}
	return -1
}

func (p Plain) EncryptAndDecode(s string) string {
	cipher := ""
	s = strings.ToLower(s)
	for _, r := range s {
		if !unicode.IsLetter(r) && !unicode.IsNumber(r) {
			continue
		}

		pos := p.getElementIndex(r)
		if pos == -1 {
			cipher += string(r)
		} else {
			cipher += string(p[25-pos])
		}
	}
	return cipher
}
func (p Plain) DivideGroup(cipher string, groups int) string {
	group := ""
	for i, c := range cipher {
		if i%groups == 0 && i != 0 {
			group += " " + string(c)
			continue
		}
		group += string(c)
	}
	return group
}

func Atbash(s string) string {
	plain := New()
	cipher := plain.EncryptAndDecode(s)
	cipher = plain.DivideGroup(cipher, 5)
	return cipher
}
