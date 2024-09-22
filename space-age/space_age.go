package space

import (
	"strings"
)

const EarthDaysSeconds = 31557600

type Planet string
type PlanetAge map[Planet]float64

func NewPlanetAge() PlanetAge {
	return PlanetAge{
		"mercury": 0.2408467, "venus": 0.61519726, "earth": 1.0, "mars": 1.8808158,
		"jupiter": 11.862615, "saturn": 29.447498, "uranus": 84.016846, "neptune": 164.79132,
	}
}

func (p PlanetAge) ConvertAgeBySeconds(seconds float64, planet Planet) float64 {
	years, ok := p[Planet(strings.ToLower(string(planet)))]
	if !ok {
		return -1.0
	}
	return seconds / (EarthDaysSeconds * years)
}

func Age(seconds float64, planet Planet) float64 {
	return NewPlanetAge().ConvertAgeBySeconds(seconds, planet)
}
