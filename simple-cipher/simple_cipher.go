package cipher

import (
	"strings"
	"unicode"
)

type Plain []rune

func newPlain() Plain {
	arr := make(Plain, 0)
	for r := 'a'; r <= 'z'; r++ {
		arr = append(arr, r)
	}
	return arr
}

type SimpleCipher struct {
	plain Plain
	mod   int
	shift int
}

func (c *SimpleCipher) Encode(text string) string {
	text = strings.ToLower(text)
	cipher := ""
	for _, r := range text {
		if !unicode.IsLetter(r) {
			continue
		}
		cipher += string(c.plain[(int(r-'a')+c.shift)%c.mod])
	}
	return cipher
}
func (c *SimpleCipher) Decode(text string) string {
	text = strings.ToLower(text)
	cipher := ""
	for _, r := range text {
		if !unicode.IsLetter(r) {
			continue
		}
		cipher += string(c.plain[(int(r-'a')-c.shift+c.mod)%c.mod])
	}
	return cipher
}

func NewCaesar() Cipher {
	p := newPlain()
	return &SimpleCipher{
		plain: p,
		mod:   len(p),
		shift: 3,
	}
}

func NewShift(distance int) Cipher {
	if distance < 0 && distance > -26 {
		distance = 26 + distance
	} else if distance <= -26 || distance >= 26 || distance == 0 {
		return nil
	}

	p := newPlain()
	return &SimpleCipher{
		shift: distance,
		mod:   len(p),
		plain: p,
	}
}

type vigenere struct {
	key string
}

func NewVigenere(key string) Cipher {
	isValidKey := func(key string) bool {
		for _, r := range key {
			if r < 'a' || r > 'z' {
				return false
			}
		}
		return len(key) > 0 && strings.Repeat("a", len(key)) != key
	}

	if !isValidKey(key) {
		return nil
	}
	return &vigenere{key: key}
}

func (v *vigenere) Encode(input string) string {
	input = v.normalizeText(input)
	key := v.normalizeKey(len(input))
	var encoded strings.Builder
	for i, r := range input {
		if !unicode.IsLetter(r) {
			continue
		}
		shift := int(key[i] - 'a')
		encoded.WriteRune(rune((int(r-'a')+shift)%26 + 'a'))
	}
	return encoded.String()
}

func (v *vigenere) Decode(input string) string {
	input = v.normalizeText(input)
	key := v.normalizeKey(len(input))
	var decoded strings.Builder
	for i, r := range input {
		if !unicode.IsLetter(r) {
			continue
		}
		shift := int(key[i] - 'a')
		decoded.WriteRune(rune((int(r-'a')-shift+26)%26 + 'a'))
	}
	return decoded.String()
}

func (v *vigenere) normalizeKey(length int) string {
	var normalized strings.Builder
	for i := 0; i < length; i++ {
		normalized.WriteByte(v.key[i%len(v.key)])
	}
	return normalized.String()
}

func (v *vigenere) normalizeText(text string) string {
	text = strings.ToLower(text)
	var normalized strings.Builder
	for _, ch := range text {
		if unicode.IsLetter(ch) {
			normalized.WriteRune(ch)
		}
	}
	return normalized.String()
}
