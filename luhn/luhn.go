package luhn

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// Using a struct (Luhn) allows better encapsulation and scalability,
// making it easier to extend the functionality (e.g., adding new methods)
// without affecting the global state or requiring major refactoring in the future.
type Luhn struct{}

// process Digits doubles the digit at every even position
func (l *Luhn) processDigits(id string) (string, error) {

	for i := len(id) - 2; i >= 0; i -= 2 {
		num, err := strconv.Atoi(string(id[i]))

		if err != nil {
			return "", fmt.Errorf("fail to convert a letter to num: pos %d, letter %v", i, id[i])
		}
		num *= 2

		if num > 9 {
			num -= 9
		}
		id = id[:i] + strconv.Itoa(num) + id[i+1:]
	}
	return id, nil
}

// sumDigit calculates the sum of all digits in the string
func (l *Luhn) sumDigit(id string) (int, error) {
	sum := 0

	for _, c := range id {
		num, err := strconv.Atoi(string(c))
		if err != nil {
			return 0, fmt.Errorf("fail to check char %v, %v ", c, err)
		}
		sum += num
	}
	return sum, nil
}

// Check verifies if the input ID is valid according to the Luhn algorithm
func (l *Luhn) Check(id string) (bool, error) {
	id = strings.Replace(id, " ", "", -1)

	if len(id) <= 1 {
		return false, errors.New("invalid id length")
	}

	process, err := l.processDigits(id)

	if err != nil {
		return false, err
	}

	sum, err := l.sumDigit(process)
	return err == nil && sum%10 == 0, err
}

// NewLuhn creates a new instance of Luhn
func NewLuhn() *Luhn {
	return &Luhn{}
}

func Valid(id string) bool {
	result, err := NewLuhn().Check(id)
	return err == nil && result
}
