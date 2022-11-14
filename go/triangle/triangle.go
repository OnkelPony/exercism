// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package triangle implements functions for determinig type of triangle.
package triangle

import "math"

// Kind represents kind of triangle.
type Kind int

const (
	// NaT means not a triangle
	NaT = iota
	// Equ means equilateral
	Equ
	// Iso means isosceles
	Iso
	// Sca means scalene
    Sca
)

// KindFromSides returns type of triangle based on the length of its sides.
func KindFromSides(a, b, c float64) Kind {
	// Write some code here to pass the test suite.
	// Then remove all the stock comments.
	// They're here to help you get started but they only clutter a finished solution.
	// If you leave them in, reviewers may protest!
	if isSideBiggerThanTwoCombined(a, b, c) || isSideNotPositive(a, b, c) || anySideNaN(a, b, c) || anySideInf(a, b, c){
		return NaT
	}
	if a == b && b == c {
		return Equ
	}
	if a == b || b == c || a == c {
		return Iso
	}
	return Sca
}

func anySideInf(a float64, b float64, c float64) bool{
	if math.IsInf(a, 1) || math.IsInf(b, 1) || math.IsInf(c, 1) {
		return true
	}
	return false
}

func anySideNaN(a float64, b float64, c float64) bool{
	if math.IsNaN(a) || math.IsNaN(b) || math.IsNaN(c){
		return true
	}
	return false
}

func isSideNotPositive(a float64, b float64, c float64) bool{
	if a <= 0 || b <= 0 || c <= 0 {
		return true
	}
	return false
}

func isSideBiggerThanTwoCombined(a float64, b float64, c float64) bool {
	if a+b < c || a+c < b || b+c < a {
		return true
	}
	return false
}
