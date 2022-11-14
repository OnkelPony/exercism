// Package secret transforms numeric code to secret handshake.
package secret

// Handshake returns string slice with "handshakes", depending on binary value of code in argument.
func Handshake(code uint) []string {
	var result []string
	if code&1 == 1 {
		result = append(result, "wink")
	}
	if code&2 == 2 {
		result = append(result, "double blink")
	}
	if code&4 == 4 {
		result = append(result, "close your eyes")
	}
	if code&8 == 8 {
		result = append(result, "jump")
	}
	if code&16 == 16 {
		result = reverse(result)
	}
	return result
}

func reverse(result []string) []string {
	for i := len(result)/2 - 1; i >= 0; i-- {
		opp := len(result) - 1 - i
		result[i], result[opp] = result[opp], result[i]
	}
	return result
}
