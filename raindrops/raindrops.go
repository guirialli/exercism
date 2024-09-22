package raindrops

import "strconv"

const (
	divisibleBy3 = "Pling"
	divisibleBy5 = "Plang"
	divisibleBy7 = "Plong"
)

type RainDrops struct {
	num   int
	sound string
}

func (r *RainDrops) Sound(dividend int) string {
	if dividend%r.num == 0 {
		return r.sound
	}
	return ""
}

func Convert(number int) string {
	rainsSound := []RainDrops{{3, divisibleBy3}, {5, divisibleBy5}, {7, divisibleBy7}}
	sound := ""
	for _, rain := range rainsSound {
		sound += rain.Sound(number)
	}

	if sound == "" {
		return strconv.Itoa(number)
	}
	return sound
}
