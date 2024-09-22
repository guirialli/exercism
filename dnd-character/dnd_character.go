package dndcharacter

import (
	"math"
	"math/rand/v2"
	"sort"
)

type Character struct {
	Strength     int
	Dexterity    int
	Constitution int
	Intelligence int
	Wisdom       int
	Charisma     int
	Hitpoints    int
}

func rollDice(d int) int {
	return rand.IntN(d) + 1
}

// Modifier calculates the ability modifier for a given ability score
func Modifier(score int) int {
	return int(math.Floor((float64(score) - 10) / 2))
}

// Ability uses randomness to generate the score for an ability
// Another solution, more like D&D is more complex than the other solution!
func Ability() int {
	dices := make([]int, 4)
	points := 0
	for i := 0; i < 4; i++ {
		dices[i] = rollDice(6)
	}
	sort.Ints(dices)
	for i := 1; i < 4; i++ {
		points += dices[i]
	}
	return points
}

// GenerateCharacter creates a new Character with random scores for abilities
func GenerateCharacter() Character {
	constitution := Ability()
	modify := Modifier(constitution)
	return Character{
		Charisma:     Ability(),
		Constitution: constitution,
		Dexterity:    Ability(),
		Hitpoints:    10 + modify,
		Intelligence: Ability(),
		Strength:     Ability(),
		Wisdom:       Ability(),
	}
}
