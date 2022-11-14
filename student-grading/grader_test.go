package main

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadGradesCsv(t *testing.T) {
	students, err := readGradesCsv("resources/grades.csv")
	if err != nil {
		log.Fatal(err)
	}

	// using testing package
	if len(students) != 30 {
		t.Errorf("Expected number of students is not same as"+
			" actual number %d", len(students))
	}

	firstStudent := Student{"Kaylen", "Johnson", "Duke University", 52, 47, 35, 38}
	assert.Equal(t, students[0], firstStudent, "First Student is incorrect")

	lastStudent := Student{"Solomon", "Hunter", "Boston University", 45, 62, 32, 58}
	assert.Equal(t, students[29], lastStudent, "Last Student is incorrect")
}

func TestGradeStudents(t *testing.T) {
	students, err := readGradesCsv("resources/grades.csv")
	if err != nil {
		log.Fatal(err)
	}

	gs := gradeStudents(students)

	// using testing package
	if len(gs) != 30 {
		t.Errorf("Expected number of students is not same as"+
			" actual number %d", len(gs))
	}
	expectedStudentScores := []float64{43, 59.25, 53, 58.25, 52.25, 50.75, 54.75, 49.25, 64.75, 43.25, 68.5, 57.75, 68.25, 66.75, 45.5, 45.75, 45.5, 58, 56, 60.25, 61, 62.5, 80.5, 53, 30.75, 57.5, 70.75, 48.5, 60.25, 49.25}
	expectedStudentGrades := []Grade{C, B, B, B, B, B, B, C, B, C, B, B, B, B, C, C, C, B, B, B, B, B, A, B, F, B, A, C, B, C}

	for i := 0; i < len(gs); i++ {
		assert.Equal(t, gs[i].finalScore, expectedStudentScores[i], "Final Scores are not same")
		assert.Equal(t, gs[i].grade, expectedStudentGrades[i], "Grades are not same")
	}
}

func TestFindOverallTopper(t *testing.T) {
	students, err := readGradesCsv("resources/grades.csv")
	if err != nil {
		log.Fatal(err)
	}
	gs := gradeStudents(students)

	actualTopper := findOverallTopper(gs)
	expectedTopper := GradedStudent{Student{"Bernard", "Wilson", "Boston University", 90, 85, 76, 71}, 80.5, A}

	assert.Equal(t, actualTopper, expectedTopper, "The topper is incorrect ")
}

func TestFindTopperPerUniversity(t *testing.T) {
	students, err := readGradesCsv("resources/grades.csv")
	if err != nil {
		log.Fatal(err)
	}
	gs := gradeStudents(students)

	universityTopper := findTopperPerUniversity(gs)
	assert.Equal(t, universityTopper["Union College"], GradedStudent{Student{"Izayah", "Hunt", "Union College", 29, 78, 41, 85}, 58.25, B})
	assert.Equal(t, universityTopper["University of California"], GradedStudent{Student{"Karina", "Shaw", "University of California", 69, 78, 56, 70}, 68.25, B})
	assert.Equal(t, universityTopper["Boston University"], GradedStudent{Student{"Bernard", "Wilson", "Boston University", 90, 85, 76, 71}, 80.5, A})
	assert.Equal(t, universityTopper["Duke University"], GradedStudent{Student{"Tamara", "Webb", "Duke University", 73, 62, 90, 58}, 70.75, A})
	assert.Equal(t, universityTopper["University of Florida"], GradedStudent{Student{"Nathan", "Gordon", "University of Florida", 53, 79, 84, 51}, 66.75, B})

}
