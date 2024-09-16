package purchase

import (
	"fmt"
	"strings"
)

// NeedsLicense determines whether a license is needed to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {
	kLower := strings.ToLower(kind)
	return kLower == "car" || kLower == "truck"
}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in lexicographical order.
func ChooseVehicle(option1, option2 string) string {
	var bestVehicle string
	if option1 > option2 {
		bestVehicle = option2
	} else {
		bestVehicle = option1
	}
	return fmt.Sprintf("%s is clearly the better choice.", bestVehicle)
}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {
	var discount float64
	if age < 3 {
		discount = 0.80
	} else if age >= 10 {
		discount = 0.50
	} else {
		discount = 0.70
	}
	return (originalPrice * discount)
}
