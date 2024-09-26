package resistorcolorduo

type resitorColor map[string]int

func NewResistorColor() resitorColor {
	return resitorColor{
		"black": 0, "brown": 1, "red": 2, "orange": 3, "yellow": 4,
		"green": 5, "blue": 6, "violet": 7, "grey": 8, "white": 9,
	}
}

func Value(colors []string) int {
	res := NewResistorColor()
	return res[colors[0]]*10 + res[colors[1]]
}
