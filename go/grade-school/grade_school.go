// Package school implements graduating school students.
package school

import (
	"sort"
)

// Grade is the base grade struct.
type Grade struct {
	level    int
	students []string
}

// School is the base school type.
type School struct {
	grades map[int][]string
}

// New returns new school.
func New() *School {
	return &School{
		grades: make(map[int][]string, 0),
	}
}

// Add adds student into grade.
func (s *School) Add(name string, level int) {
	s.grades[level] = append(s.grades[level], name)
}

// Grade returns students for given grade.
func (s *School) Grade(level int) []string {
	return s.grades[level]
}

// Enrollment enroll all students.
func (s *School) Enrollment() []Grade {
	result := make([]Grade, len(s.grades))

	keys := make([]int, len(s.grades))
	i := 0
	for k := range s.grades {
		keys[i] = k
		i++
	}

	sort.Ints(keys)

	for i, k := range keys {
		v := s.grades[k]
		sort.Strings(v)
		result[i] = Grade{
			level:    k,
			students: v,
		}
	}

	return result
}
