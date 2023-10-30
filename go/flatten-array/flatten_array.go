package flatten

// Flatten expects an (nested) slice / array and returns the flattened slice
func Flatten(arr interface{}) []interface{} {
	result := []interface{}{}

	switch elem := arr.(type) {
	case []interface{}:
		for _, v := range elem {
			result = append(result, Flatten(v)...)
		}
	case interface{}:
		result = append(result, elem)
	}
	return result
}
