package collatzconjecture

import (
	"errors"
)

func collatz(n int, step int) int {
	if n == 1 {
		return step
	} else if n%2 == 0 {
		return collatz(n/2, step+1)
	}
	return collatz(3*n+1, step+1)
}

func CollatzConjecture(n int) (int, error) {
	if n <= 0 {
		return 0, errors.New("n must be positive integer and more than zero")
	}
	return collatz(n, 0), nil

}
