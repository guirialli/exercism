package prime

func Factors(n int64) []int64 {
	var factor int64 = 2
	primes := make([]int64, 0)
	for n >= int64(factor) {
		if n%factor == 0 {
			primes = append(primes, factor)
			n = n / factor
			continue
		}
		factor++
	}
	return primes
}
