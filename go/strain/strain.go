// Package strain implements functions, which return slices based on predicate in argument
package strain

// Ints is list of integers
type Ints []int

// Keep returns Ints for which members predicate in argument is true
func (i Ints) Keep(pred func(int) bool) Ints {
	var result Ints
	for _, number := range i {
		if pred(number) {
			result = append(result, number)
		}
	}
	return result
}

// Discard returns Ints for which members predicate in argument is false
func (i Ints) Discard(pred func(int) bool) interface{} {
	var result Ints
	for _, number := range i {
		if !pred(number) {
			result = append(result, number)
		}
	}
	return result
}

// Lists is a list of list of integers
type Lists [][]int

// Keep returns Lists for which members predicate in argument is true
func (l Lists) Keep(has5 func(l []int) bool) interface{} {
	var result Lists
	for _, intList := range l {
		if has5(intList) {
			result = append(result, intList)
		}
	}
	return result
}

// Strings is a list of strings
type Strings []string

// Keep returns Strings for which members predicate in argument is true
func (s Strings) Keep(zword func(s string) bool) interface{} {
	var result Strings
	for _, word := range s {
		if zword(word) {
			result = append(result, word)
		}
	}
	return result
}
