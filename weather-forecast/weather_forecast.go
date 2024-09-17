// Package weather provides tools for Goblinocus weather operators to get the current weather
// forecast for a specific location. It allows them to store and retrieve the weather conditions
// for a given city, helping them keep track of current conditions in their region.
package weather

// CurrentCondition stores the most recent weather condition as a string.
// It helps track the state of the weather for a particular location.
var CurrentCondition string

// CurrentLocation stores the name of the location (city) for which the weather condition is recorded.
// It allows operators to associate the recorded weather condition with a specific place.
var CurrentLocation string

// Forecast takes a city and a weather condition as inputs and returns a message
// showing the current weather condition for that city.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
