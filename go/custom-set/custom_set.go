package stringset

import "fmt"

// Implement Set as a collection of unique string values.
//
// For Set.String, use '{' and '}', output elements as double-quoted strings
// safely escaped with Go syntax, and use a comma and a single space between
// elements. For example, a set with 2 elements, "a" and "b", should be formatted as {"a", "b"}.
// Format the empty set as {}.

// Define the Set type here.
type Set map[string]struct{}

func New() Set {
	return make(Set)
}

func NewFromSlice(l []string) Set {
	result := New()
	for _, s := range l {
		result.Add(s)
	}
	return result
}

func (s Set) String() string {
	var i int
	result := "{"
	for k := range s {
		i++
		result += fmt.Sprintf("\"%v\"", k)
		if i < len(s) {
			result += ", "
		}
	}
	result += "}"
	return result
}

func (s Set) IsEmpty() bool {
	return len(s) == 0
}

func (s Set) Has(elem string) bool {
	_, ok := s[elem]
	return ok
}

func (s Set) Add(elem string) {
	s[elem] = struct{}{}
}

func Subset(s1, s2 Set) bool {
	for k := range s1 {
		if !s2.Has(k) {
			return false
		}
	}
	return true
}

func Disjoint(s1, s2 Set) bool {
	panic("Please implement the Disjoint function")
}

func Equal(s1, s2 Set) bool {
	return Subset(s1, s2) && len(s1) == len(s2)
}

func Intersection(s1, s2 Set) Set {
	panic("Please implement the Intersection function")
}

func Difference(s1, s2 Set) Set {
	panic("Please implement the Difference function")
}

func Union(s1, s2 Set) Set {
	panic("Please implement the Union function")
}
