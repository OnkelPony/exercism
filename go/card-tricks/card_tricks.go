package cards

// GetItem retrieves an item from a slice at given position. The second return value indicates whether
// the given index exists in the slice or not.
func GetItem(slice []int, index int) (int, bool) {
	if index >= 0 && index < len(slice) {
		return slice[index], true
	} else {
		return 0, false
	}
}

// SetItem writes an item to a slice at given position overwriting an existing value.
// If the index is out of range the value needs to be appended.
func SetItem(slice []int, index, value int) []int {
	_, inSlice := GetItem(slice, index)
	if !inSlice {
		return append(slice, value)
	}
	slice[index] = value
	return slice
}

// PrefilledSlice creates a slice of given length and prefills it with the given value.
func PrefilledSlice(value, length int) []int {
	if length < 1 {
		return []int{}
	}
	result := make([]int, length)
	for i := 0; i < length; i++ {
		result[i] = value
	}
	return result
}

// RemoveItem removes an item from a slice by modifying the existing slice.
func RemoveItem(slice []int, index int) []int {
	_, inSlice := GetItem(slice, index)
	if !inSlice {
		return slice
	}
	return append(slice[:index], slice[index+1:]...)
}
