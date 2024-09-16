package cars

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	// Convert productionRate to Float
	productionRateFloat := float64(productionRate)
	return productionRateFloat * (successRate / 100)
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	carsPerHours := CalculateWorkingCarsPerHour(productionRate, successRate)
	return int(carsPerHours / 60)
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
	/*
		Each car normally costs $10,000 to produce individually, regardless of whether it is successful or not.
		But with a bit of planning, 10 cars can be produced together for $95,000.
	*/
	produceUnit := 10000
	produceGroup10 := 95000

	num_groups := carsCount / 10
	remaning := carsCount % 10
	return uint((produceGroup10 * num_groups) + (produceUnit * remaning))
}
