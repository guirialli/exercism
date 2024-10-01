package prime

import (
	"fmt"
	"math"
)

func isNumberPrime(n int) bool {
	for i := 2; i < int(math.Sqrt(float64(n)))+1; i++ {
		if n%i == 0 && i != n {
			return false
		}
	}
	return true
}

func Nth(nth int) (int, error) {
	if nth < 1 {
		return 0, fmt.Errorf("nth %d: prime number must be greater than zero", nth)
	}
	lastPrime := 2
	for nth > 1 {
		lastPrime++
		if isNumberPrime(lastPrime) {
			nth--
		}
	}
	return lastPrime, nil
}
