package isbn

import (
	"strconv"
	"strings"
	"unicode"
)

func IsValidISBN(isbn string) bool {
	isbn = strings.ReplaceAll(isbn, "-", "")
	length := len(isbn)
	if isbn == "" || length != 10 || (isbn[length-1] != 'X' && unicode.IsLetter(rune(isbn[length-1]))) {
		return false
	}

	sum := 0
	for i, v := range isbn {
		n, err := strconv.Atoi(string(v))
		if err != nil {
			if i != 9 {
				return false
			}
			n = 10
		}
		sum += n * (10 - i)
	}
	return sum%11 == 0
}
