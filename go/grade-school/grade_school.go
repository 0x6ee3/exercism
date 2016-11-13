package school

import "sort"

// Grade ..
type Grade struct {
	year     int
	students []string
}

// School ..
type School map[int]Grade

// New ..
func New() *School {
	return &School{}
}

// Add ..
func (s *School) Add(name string, year int) {
	if _, ok := (*s)[year]; !ok {
		(*s)[year] = Grade{year: year, students: []string{}}
	}
	students := (*s)[year].students
	students = append(students, name)
	(*s)[year] = Grade{year: year, students: students}
}

// Grade ..
func (s *School) Grade(year int) []string {
	return (*s)[year].students
}

// Enrollment ..
func (s *School) Enrollment() []Grade {
	var years []int
	for year := range *s {
		years = append(years, year)
	}
	sort.Ints(years)

	var grades []Grade
	for _, year := range years {
		grade := (*s)[year]
		sort.Strings(grade.students)
		grades = append(grades, grade)
	}

	return grades
}
