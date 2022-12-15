// Package darts implements method to calculate in which part of the target dart landed
package darts

import "math"

const (
	one  = 1
	five = 5
	ten  = 10
)

// Score returns amount of points based on cartesian coordinates of the dart tossed
func Score(x float64, y float64) int {
	r := math.Sqrt(x * x + y * y)
	if r <= one {
		return ten
	}
	if r <= five {
		return five
	}
	if r <= ten {
		return one
	}
	return 0
}
