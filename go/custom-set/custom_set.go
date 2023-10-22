package stringset

import (
	"strings"
)

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
	var out strings.Builder
	out.WriteByte('{')
	len := len(s)
	for k := range s {
		out.WriteByte('"')
		out.WriteString(k)
		out.WriteByte('"')
		if len--; len > 0 {
			out.WriteString(", ")
		}
	}
	out.WriteByte('}')
	return out.String()
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
	return len(Intersection(s1, s2)) == 0
}

func Equal(s1, s2 Set) bool {
	return Subset(s1, s2) && len(s1) == len(s2)
}

func Intersection(s1, s2 Set) Set {
	result := New()
	for k := range s1 {
		if s2.Has(k) {
			result.Add(k)
		}
	}
	return result
}

func Difference(s1, s2 Set) Set {
	result := New()
	for s := range s1 {
		if !s2.Has(s) {
			result.Add(s)
		}
	}
	return result
}

func Union(s1, s2 Set) Set {
	for s := range s1 {
		s2.Add(s)
	}
	return s2
}
