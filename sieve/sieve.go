package sieve

import (
	"math"
)

func isPrime(n int) bool {
	for i := 2; i <= int(math.Sqrt(float64(n)))+1; i++ {
		if n%i == 0 && i != n {
			return false
		}
	}
	return true
}

func Sieve(limit int) []int {
	primes := make([]int, 0)
	sieves := make(map[int]bool)
	for i := 2; i <= limit; i++ {
		if isPrime(i) {
			primes = append(primes, i)
			for m := i * i; m <= limit; m *= i {
				sieves[m] = true
			}
		}
	}
	return primes
}
