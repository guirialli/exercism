package hamming

import "errors"

func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, errors.New("cells are not the same length")
	}
	hamming := 0
	for i, _ := range a {
		if a[i] != b[i] {
			hamming++
		}

	}
	return hamming, nil
}
