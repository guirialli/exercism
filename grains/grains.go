package grains

import (
	"fmt"
)

const total uint64 = (1 << 64) - 1

func Square(number int) (uint64, error) {
	if number <= 0 || number > 64 {
		return 0, fmt.Errorf("number %d is out of range, n > 0 && n <= 64", number)
	}
	return 1 << uint(number-1), nil
}

func Total() uint64 {
	return total
}
