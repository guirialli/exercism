package armstrong

import (
	"math"
	"strconv"
)

func IsNumber(n int) bool {
	strN := strconv.Itoa(n)
	sum := 0
	length := float64(len(strN))
	for _, c := range strN {
		n, _ := strconv.Atoi(string(c))
		sum += int(math.Pow(float64(n), length))
	}
	return n == sum
}
