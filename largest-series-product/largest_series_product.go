package lsproduct

import (
	"fmt"
	"strconv"
)

func LargestSeriesProduct(digits string, span int) (int64, error) {
	if len(digits) < span || span < 0 {
		return 0, fmt.Errorf("span must be smaller than string length")
	}
	var total int64 = 0
	for i := 0; i+span <= len(digits); i++ {
		var subTotal int64 = 1
		for _, c := range digits[i : span+i] {
			digit, err := strconv.Atoi(string(c))
			if err != nil {
				return 0, err
			}
			subTotal *= int64(digit)
		}
		if subTotal > total {
			total = subTotal
		}
	}

	return total, nil
}
