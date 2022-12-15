// Package sieve implements prime finding algorithm Sieve of Eratosthenes.
package sieve

import (
	"math"
	"sort"
)

// Sieve returns all primes up to the limit in argument.
func Sieve(limit int) []int {
	const firstPrime = 2
	const secondPrime = 3
	sieve := make(map[int]bool)
	var result []int
	if limit < firstPrime {
		return result
	}
	result = append(result, firstPrime)
	for i := secondPrime; i <= limit; i += 2 {
		sieve[i] = true
	}
	for j := secondPrime; j <= int(math.Sqrt(float64(limit))); j += 2 {
		if sieve[j] == true {
			for k := j * j; k <= limit; k = k + j {
				sieve[k] = false
			}
		}
	}
	for l := range sieve {
		if sieve[l] == true {
			result = append(result, l)
		}
	}
	sort.Ints(result)
	return result
}
