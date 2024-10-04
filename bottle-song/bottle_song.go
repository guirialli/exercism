package bottlesong

import (
	"fmt"
	"strings"
)

type bottler struct{}

var numLetter = map[int]string{0: "No", 1: "One", 2: "Two", 3: "Three", 4: "Four", 5: "Five", 6: "Six", 7: "Seven", 8: "Eight", 9: "Nine", 10: "Ten"}

func (b *bottler) singularPlural(verse int) string {
	if verse == 1 {
		return "bottle"
	}
	return "bottles"
}

func (b *bottler) getVerses(verse int) []string {
	versers := make([]string, 4)
	versers[0] = fmt.Sprintf("%s green %s hanging on the wall,", numLetter[verse], b.singularPlural(verse))
	versers[1] = fmt.Sprintf("%s green %s hanging on the wall,", numLetter[verse], b.singularPlural(verse))
	versers[2] = "And if one green bottle should accidentally fall,"
	versers[3] = fmt.Sprintf("There'll be %s green %s hanging on the wall.", strings.ToLower(numLetter[verse-1]), b.singularPlural(verse-1))
	return versers
}

func NewBottler() *bottler {
	return &bottler{}
}

func Recite(startBottles, takeDown int) []string {
	versers := make([]string, 0)
	b := NewBottler()
	for takeDown > 0 {
		versers = append(versers, b.getVerses(startBottles)...)
		startBottles -= 1
		takeDown -= 1
		if takeDown > 0 {
			versers = append(versers, "")
		}
	}
	return versers
}
