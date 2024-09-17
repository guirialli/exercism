package gross

// Units stores the Gross Store unit measurements.
/*
Unit	Score
quarter_of_a_dozen	3
half_of_a_dozen	6
dozen	12
small_gross	120
gross	144
great_gross	1728
*/
func Units() map[string]int {
	units := make(map[string]int)
	units["quarter_of_a_dozen"] = 3
	units["half_of_a_dozen"] = 6
	units["dozen"] = 12
	units["small_gross"] = 120
	units["gross"] = 144
	units["great_gross"] = 1728
	return units
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	return make(map[string]int)
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	value, ok := units[unit]
	if ok == false {
		return false
	}

	valueBill, _ := bill[item]
	bill[item] = valueBill + value
	return true

}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	unitValue, unitExist := units[unit]
	itemBill, itemBillExist := bill[item]
	newQuantity := itemBill - unitValue

	if (!unitExist || !itemBillExist) || (itemBill < 0 || newQuantity < 0) {
		return false
	} else if newQuantity == 0 {
		delete(bill, item)
	} else {
		bill[item] = newQuantity
	}

	return true
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	i, ok := bill[item]
	return i, ok
}
