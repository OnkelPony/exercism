// Package summultiples implements method, which sums all multiples.
package summultiples

// SumMultiples returns sum of all numbers, divisible by any divisor up to the limit in arguments.
func SumMultiples(limit int, divisors ...int) int {
	if len(divisors) == 0 {
		return 0
	}
	result := 0
	for i := 1; i < limit; i++ {
		if isDivisible(divisors, i) {
			result += i
		}
	}
	return result
}

// isDivisible returns true for each number, divisible by one of divisors in argument.
func isDivisible(divisors []int, i int) bool {
	for _, number := range divisors {
		if number == 0 {
			continue
		}
		if i%number == 0 {
			return true
		}
	}
	return false
}
