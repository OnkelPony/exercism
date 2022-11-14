// Package accumulate implements method, which returns slice argument, modified by func argument
package accumulate

// Accumulate converts slice using function both entered as arguments and returns result of it.
func Accumulate(given []string, converter func(string) string) []string {
	givenLength := len(given)
	if givenLength == 0 {
		return given
	}
	result := make([]string, givenLength)
	for i, item := range given {
		result[i] = converter(item)
	}
	return result
}

