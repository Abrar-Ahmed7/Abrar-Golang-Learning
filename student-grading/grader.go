package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

type Grade string

const (
	A Grade = "A"
	B Grade = "B"
	C Grade = "C"
	F Grade = "F"
)

type Student struct {
	firstName, lastName, university string
	test1, test2, test3, test4      int
}

type GradedStudent struct {
	Student
	finalScore float64
	grade      Grade
}

func main() {
	rows, err := readGradesCsv("resources/grades.csv")
	if err != nil {
		log.Fatal(err)
	}

	// Parsing all Students
	fmt.Print("\n *******All Students*******")
	for _, student := range rows {
		fmt.Println(student.string())
	}

	// Calculating Grades
	fmt.Print("******Grades of all the students******\n")
	for _, gs := range gradeStudents(rows) {
		fmt.Println(gs.string())
	}

	gs := gradeStudents(rows)

	// Finding Over All Topper
	fmt.Print("\n******Over All Topper******\n")
	fmt.Println(findOverallTopper(gs).string())

	// Finding Topper per University
	fmt.Print("\n******Topper of Each University******\n")
	for uni, topper := range findTopperPerUniversity(gs) {
		fmt.Println(uni, "-", topper.string())
	}

}

func readGradesCsv(path string) ([]Student, error) {
	// Skip first row (line)
	csvFile, err := os.Open("resources/grades.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer csvFile.Close()
	row1, err := bufio.NewReader(csvFile).ReadSlice('\n')
	if err != nil {
		return nil, err
	}
	_, err = csvFile.Seek(int64(len(row1)), io.SeekStart)
	if err != nil {
		return nil, err
	}
	var students []Student
	// Read remaining rows
	r := csv.NewReader(csvFile)
	for {
		line, err := r.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		students = append(students, Student{
			line[0],
			line[1],
			line[2],
			convertStrToInt(line[3]),
			convertStrToInt(line[4]),
			convertStrToInt(line[5]),
			convertStrToInt(line[6]),
		})
	}
	return students, nil
}

func gradeStudents(students []Student) []GradedStudent {
	var gradeComputedStudents []GradedStudent
	for _, student := range students {
		finalScore := calculateFinalScore(student)
		allTestScoresgradeComputedStudent := GradedStudent{
			student,
			finalScore,
			calculateGrade(finalScore),
		}
		gradeComputedStudents = append(gradeComputedStudents, allTestScoresgradeComputedStudent)
	}
	return gradeComputedStudents
}

func findOverallTopper(gradedStudents []GradedStudent) GradedStudent {
	var topScore float64
	var topGradedStudentIndex int
	for i, gradedStudent := range gradedStudents {
		if gradedStudent.finalScore > topScore {
			topScore = gradedStudent.finalScore
			topGradedStudentIndex = i
		}
	}
	return gradedStudents[topGradedStudentIndex]
}

func findTopperPerUniversity(gradedStudents []GradedStudent) map[string]GradedStudent {
	var topperPerUniversity = make(map[string]GradedStudent)
	for _, gs := range gradedStudents {
		// if gs, ok := topperPerUniversity[gs.university]; ok {
		// 	topperPerUniversity[gs.university] = gs
		// }
		if topperPerUniversity[gs.university].finalScore < gs.finalScore {
			topperPerUniversity[gs.university] = gs
		}
	}
	return topperPerUniversity
}

func convertStrToInt(str string) int {
	var num, err = strconv.Atoi(str)
	if err != nil {
		log.Fatal(err)
	}
	return num
}

func calculateFinalScore(s Student) float64 {
	return float64(s.test1+s.test2+s.test3+s.test4) / 4
}

func calculateGrade(finalScore float64) Grade {
	if finalScore < 35 {
		return F
	} else if finalScore >= 35 && finalScore < 50 {
		return C
	} else if finalScore >= 50 && finalScore < 70 {
		return B
	} else {
		return A
	}
}

func (s Student) string() string {
	return fmt.Sprintf("%v %v %v %v %v %v %v", s.firstName, s.lastName, s.university, s.test1, s.test2, s.test3, s.test4)
}

func (gs GradedStudent) string() string {
	return fmt.Sprintf("%v %v %v", gs.Student.string(), gs.finalScore, gs.grade)
}
