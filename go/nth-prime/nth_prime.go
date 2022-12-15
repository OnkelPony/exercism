// Package prime implements method, which returns nth prime number.
package prime

import "math"

// Nth returns nth prime number based on argument.
func Nth(n int) (int, bool) {
	if n == 0 {
		return 0, false
	}
	if n == 1 {
		return 2, true
	}
	if n == 2 {
		return 3, true
	}
	primes := []int{3}
	for i := 5; ; i += 2 {
		isPrime := true
		sqrti := math.Sqrt(float64(i))
		for _, prime := range primes {
			if i%prime == 0 {
				isPrime = false
				break
			}
			if float64(prime) > sqrti {
				break
			}
		}
		if isPrime {
			primes = append(primes, i)
			// We didn't count 2 as prime to avoid unnecessary testing divisibility by it.
			if len(primes)+1 == n {
				return i, true
			}
		}
	}
	return -1, false
}
