package resistorcolor

import (
	"fmt"
)

// Color is a map that associates resistor color names with their corresponding numeric values.
// The same solution, but I use slices
type Color []string

// ResistorColor holds the Color map and provides methods to work with resistor colors.
type ResistorColor struct {
	colors Color
}

func NewResistorColor() *ResistorColor {
	return &ResistorColor{
		colors: Color{
			"black", "brown", "red", "orange", "yellow",
			"green", "blue", "violet", "grey", "white",
		},
	}
}

// GetOnlyColorsName returns a slice containing the names of all resistor colors.
func (c *ResistorColor) GetOnlyColorsName() []string {
	return c.colors
}

// GetColor returns the numeric value associated with a resistor color name.
func (c *ResistorColor) GetColor(color string) (int, error) {
	for i, colorName := range c.colors {
		if colorName == color {
			return i, nil
		}
	}
	return -1, fmt.Errorf("color %s not found", color)
}

// Colors returns the list of all colors.
func Colors() []string {
	return NewResistorColor().GetOnlyColorsName()
}

// ColorCode returns the resistance value of the given color.
func ColorCode(color string) int {
	value, err := NewResistorColor().GetColor(color)
	if err != nil {
		panic(err)
	}
	return value
}
