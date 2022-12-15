// Package pythagorean implements methods, which create pythagorean triplets.
package pythagorean

// Triplet is pythagorean triplet.
type Triplet [3]int

// Range returns pythagorean triplets, whose side length is in interval in arguments.
func Range(min int, max int) []Triplet {
	var result []Triplet
	for c := min; c <= max; c++ {
		for b := min; b <= c; b++ {
			for a := min; a <= b; a++ {
				if a*a+b*b == c*c {
					result = append(result, Triplet{a, b, c})
				}
			}
		}
	}
	return result
}

// Sum returns pythagorean triplets, whose side sum is in argument.
func Sum(sum int) []Triplet {
	var result []Triplet
	for a := 1; a <= sum/3; a++ {
		for b := a; b <= 2*sum/3; b++ {
			c := sum - (a + b)
			if a*a+b*b == c*c {
				result = append(result, Triplet{a, b, c})
			}
		}
	}

	return result
}
