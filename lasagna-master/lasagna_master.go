package lasagna

import "strings"

// TODO: define the 'PreparationTime()' function
// PreparationTime receives a slice of strings and the time; by default, time is 2.
func PreparationTime(layers []string, time int) int {
	qtdLayers := len(layers)
	if time <= 0 {
		return qtdLayers * 2
	}
	return qtdLayers * time
}

// TODO: define the 'Quantities()' function
// Calculate the quantities of sauce and noodles needed. The default values are: sauce: 0.2 & noodles: 50.
func Quantities(list []string) (int, float64) {
	var sauce float64
	var noodles int
	for _, item := range list {
		iLower := strings.ToLower(item)
		switch iLower {
		case "sauce":
			sauce += 0.2
		case "noodles":
			noodles += 50
		}
	}

	return noodles, sauce
}

// TODO: define the 'AddSecretIngredient()' function
// This changes the last value of myList with the last value of friendsList.
func AddSecretIngredient(friendsList []string, myList []string) {
	myList[len(myList)-1] = friendsList[len(friendsList)-1]
}

// TODO: define the 'ScaleRecipe()' function
// This function converts the portion size to the quantity informed.
func ScaleRecipe(quantities []float64, portions int) []float64 {
	amount := float64(portions) / 2
	newQuantities := make([]float64, len(quantities))

	for i, p := range quantities {
		newQuantities[i] = amount * p
	}
	return newQuantities
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
