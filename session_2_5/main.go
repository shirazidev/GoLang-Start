package main

import "fmt"

type Student struct {
	name   string
	ID     int
	scores []int
}

func main() {
	// method

	// create an instance of the student type
	s := Student{ID: 1, name: "AmirHossein Shirazi", scores: []int{100, 89, 75, 95, 50}}
	s.printStudent()
	fmt.Println("The student's average score is: ", s.studentAvgScore())
	fmt.Println("The student's eligibility status: ", s.isStudentEligible())
}

func (s Student) printStudent() {
	fmt.Printf("Student name is %s and SID is %d.\n", s.name, s.ID)
}

func (s Student) studentAvgScore() float64 {
	sum := 0
	for _, score := range s.scores {
		sum += score
	}
	return float64(sum / len(s.scores))
}

func (s Student) isStudentEligible() bool {
	if s.studentAvgScore() >= 85 {
		return true
	}
	return false
}

//func printStudent(s Student) {
//	fmt.Printf("Student name is %s and SID is %d.\n", s.name, s.ID)
//}
//
//func studentAvgScore(s Student) float64 {
//	sum := 0
//	for _, score := range s.scores {
//		sum += score
//	}
//	return float64(sum / len(s.scores))
//}
//
//func isStudentEligible(s Student) bool {
//	if studentAvgScore(s) >= 85 {
//		return true
//	}
//	return false
//}
