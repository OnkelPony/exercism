package resistorcolortrio

import "strconv"

var codes = map[string]int{
	"black":  0,
	"brown":  1,
	"red":    2,
	"orange": 3,
	"yellow": 4,
	"green":  5,
	"blue":   6,
	"violet": 7,
	"grey":   8,
	"white":  9,
}

// Label describes the resistance value given the colors of a resistor.
// The label is a string with a resistance value with an unit appended
// (e.g. "33 ohms", "470 kiloohms").
func Label(colors []string) string {
	value := (10*codes[colors[0]] + codes[colors[1]]) * tenPow(codes[colors[2]])
	return strconv.Itoa(value) + " ohms"
}

func tenPow(n int) int {
	result := 1
	for i := 0; i < n; i++ {
		result *= 10
	}
	return result
}