package romannumerals

import (
	"errors"
)

type RomanNumeral struct {
	value  int
	letter string
}

var romanNumerals = []RomanNumeral{
	{1000, "M"}, {900, "CM"},
	{500, "D"}, {400, "CD"},
	{100, "C"}, {90, "XC"},
	{50, "L"}, {40, "XL"},
	{10, "X"}, {9, "IX"},
	{5, "V"}, {4, "IV"},
	{1, "I"},
}

func ToRomanNumeral(input int) (string, error) {
	if input >= 4000 || input < 1 {
		return "", errors.New("Invalid input")
	}
	roman := ""
	for _, num := range romanNumerals {
		for num.value <= input {
			roman += num.letter
			input -= num.value
		}
	}
	return roman, nil
}
