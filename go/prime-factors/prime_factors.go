// Package prime implements method, which returns prime factors of given number.
package prime

import "math"

// Factors returns slice of primes, which make number in argument by multiplying them together.
func Factors(input int64) []int64 {
	result := make([]int64, 0)
	var divisor int64
	for divisor = 2; ; divisor++ {
		if divisor%2 == 0 && divisor > 2 {
			continue
		}
		for {
			if input%divisor == 0 {
				result = append(result, divisor)
				input /= divisor
				continue
			}
			break
		}
		if input == 1 {
			break
		}
		if divisor >= int64(math.Sqrt(float64(input))) {
			result = append(result, input)
			break
		}
	}
	return result
}
