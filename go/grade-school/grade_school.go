package school

import "sort"

// Define the Grade and School types here.
type Grade struct {
	level    int
	students sort.StringSlice
}

type School struct {
	grades   []Grade
	students map[int][]string
}

// New returns a new instance of a School.
func New() *School {
    return &School{
        grades:   make([]Grade, 0),
        students: make(map[int][]string),
    }
}

// Add adds a student to the appropriate grade level.
func (s *School) Add(student string, grade int) {
    s.students[grade] = append(s.students[grade], student)
}

// Grade returns the students in the specified grade level.
func (s *School) Grade(level int) []string {
    return s.students[level]
}

// Enrollment returns the list of students and their grades in the school.
// The grades are sorted ascending and so are students in each grade.
func (s *School) Enrollment() []Grade {
	for grade, students := range s.students {
		s.grades = append(s.grades, Grade{grade, students})
	}
	for _, grade := range s.grades {
		grade.students.Sort()
	}
	sort.Slice(s.grades, func(i, j int) bool {
		return s.grades[i].level < s.grades[j].level
	})
	return s.grades
}
