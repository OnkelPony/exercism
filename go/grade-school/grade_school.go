package school

// Define the Grade and School types here.
type Grade struct {
	level    int
	students []string
}

type School struct {
	grades   []Grade
	students map[int][]string
}

func New() *School {
	return &School{grades: make([]Grade, 0), students: make(map[int][]string)}
}

func (s *School) Add(student string, grade int) {
	// if _, ok := s.students[grade];!ok {
	// 	s.students[grade]
	s.students[grade] = append(s.students[grade], student)
}

func (s *School) Grade(level int) []string {
return s.students[level]}

func (s *School) Enrollment() []Grade {
	if len(s.students) == 0 {
		return []Grade{}
	}
	for grade, students := range s.students {
		s.grades = append(s.grades, Grade{grade, students})
	}
	return s.grades
}
