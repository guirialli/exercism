package anagram

import (
	"sort"
	"strings"
)

type Table map[rune]int

func Detect(subject string, candidates []string) []string {
	filter := make([]string, 0)
	subject = strings.ToLower(subject)
	for _, c := range candidates {
		cLower := strings.ToLower(c)
		if cLower == subject || sortString(subject) != sortString(cLower) {
			continue
		}
		filter = append(filter, c)
	}
	return filter
}

func sortString(word string) string {
	chars := strings.Split(word, "")
	sort.Strings(chars)
	return strings.Join(chars, "")
}
