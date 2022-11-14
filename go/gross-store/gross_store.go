package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	result := map[string]int{
		"quarter_of_a_dozen": 3,
		"half_of_a_dozen":    6,
		"dozen":              12,
		"small_gross":        120,
		"gross":              144,
		"great_gross":        1728,
	}
	return result
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	result := make(map[string]int)
	return result
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	_, exists := units[unit]
	if exists {
		bill[item] += units[unit]
		return true
	}
	return false
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	items, onBill := bill[item]
	quantity, onUnits := units[unit]
	if !onBill || !onUnits || items < quantity {
		return false
	}
	if items == quantity {
		delete(bill, item)
		return true
	}
	bill[item] -= quantity
	return true
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	quantity, ok := bill[item]
	return quantity, ok
}
