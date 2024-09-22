package etl

import "strings"

func Transform(in map[int][]string) map[string]int {
	m := make(map[string]int)
	for k, v := range in {
		for _, letter := range v {
			m[strings.ToLower(letter)] = k
		}
	}
	return m
}
