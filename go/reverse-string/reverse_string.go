// Package reverse implements method, which returns reversed string in argument
package reverse

//Reverse returns string in argument reversed
func Reverse(toReverse string) string {
	if toReverse == "" {
		return toReverse
	}
	reversed := ""
	toReverseRunes := []rune(toReverse)
	for i := len(toReverseRunes) - 1; i >= 0; i-- {
		reversed += string(toReverseRunes[i])
	}
	return reversed
}
