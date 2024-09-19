package thefarm

import (
	"errors"
	"fmt"
)

// Version educative, as this could be a simple function, not a struct! Proceed to iterator 1.
type InvalidCows struct {
	cows    int
	message string
}

func (e InvalidCows) Error() string {
	return fmt.Sprintf("%d cows are invalid: %s", e.cows, e.message)
}

// TODO: define the 'DivideFood' function
func DivideFood(calculator FodderCalculator, cows int) (float64, error) {
	amount, err := calculator.FodderAmount(cows)
	if err != nil {
		return 0, err
	}
	factor, err := calculator.FatteningFactor()
	if err != nil {
		return 0, err
	}

	return (amount / float64(cows)) * factor, nil
}

func ValidateInputAndDivideFood(calculator FodderCalculator, number int) (float64, error) {
	if number <= 0 {
		return 0, errors.New("invalid number of cows")
	}
	return DivideFood(calculator, number)
}

// TODO: define the 'ValidateNumberOfCows' function
func ValidateNumberOfCows(cows int) error {
	if cows < 0 {
		return &InvalidCows{cows, "there are no negative cows"}
	} else if cows == 0 {
		return &InvalidCows{cows, "no cows don't need food"}
	}
	return nil
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
