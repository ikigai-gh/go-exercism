package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	measuraments := map[string]int{
		"quarter_of_a_dozen": 3,
		"half_of_a_dozen":    6,
		"dozen":              12,
		"small_gross":        120,
		"gross":              144,
		"great_gross":        1728,
	}

	return measuraments
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	newBill := make(map[string]int)

	return newBill
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	_, ok := units[unit]

	if !ok {
		return false
	}

	_, ok = bill[item]

	if ok {
		bill[item] += units[unit]
		return true
	} else {
		bill[item] = units[unit]
	}

	return true
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	_, ok := bill[item]

	if !ok {
		return false
	}

	_, ok = units[unit]

	if !ok {
		return false
	}

	quantity := bill[item] - units[unit]

	if quantity < 0 {
		return false
	}

	if quantity == 0 {
		delete(bill, item)
		return true
	}

	bill[item] -= units[unit]

	return true
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	_, ok := bill[item]

	if !ok {
		return 0, false
	}

	return bill[item], true
}
