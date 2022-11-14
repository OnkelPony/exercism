// Package diffsquares implements methods, which count square of sums, sum of squares and their difference
package diffsquares

// SquareOfSum returns square of sums of the first n integers
func SquareOfSum(n int) int {
	return (n*n + n) * (n*n + n) / 4
}

// SumOfSquares returns sum of squares of the first n integers
func SumOfSquares(n int) int {
	return (2*n*n*n + 3*n*n + n) / 6
}

// Difference returns difference between sum of squares and square of sums of the first n integers
func Difference(n int) int {
	return (3*n*n + 2*n) * (n*n - 1) / 12
}
