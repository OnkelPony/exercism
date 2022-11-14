// Package lsproduct implements method, which decides, which numbers have largest series product.
package lsproduct

import (
	"errors"
	"strconv"
	"unicode"
)

// LargestSeriesProduct returns largest series product from digits of span from argument.
func LargestSeriesProduct(digits string, span int) (int64, error) {
	var maxProduct int64
	digitsLength := len(digits)
	if span < 0 {
		return -1, errors.New("span must be greater than zero")
	}
	if span > digitsLength {
		return -1, errors.New("span can't be bigger than digits")
	}
	for _, digit := range digits {
		if !unicode.IsDigit(digit) {
			return -1, errors.New("only digits allowed")
		}
	}
	cycles := digitsLength - span
	for i := 0; i <= cycles; i++ {
		spanDigits := digits[i : i+span]
		product := makeProduct(spanDigits)
		if product > maxProduct {
			maxProduct = product
		}
	}
	return maxProduct, nil
}

// makeProduct returns product of all digits in argument
func makeProduct(digits string) int64 {
	var result int64 = 1
	for _, i := range digits {
		factor, _ := strconv.Atoi(string(i))
		result *= int64(factor)
	}
	return result
}
