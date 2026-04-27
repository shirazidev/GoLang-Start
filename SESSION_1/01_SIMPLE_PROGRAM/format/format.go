package format

import (
	"fmt"
	"simple/student"
)

func Format(p student.Student) {
	fmt.Printf("\nstudent id: %d\n National Id: %s\n Name: %s\n LastName: %s\n "+
		"Age: %v\n", p.StudentID, p.NationalID, p.FirstName, p.LastName, p.Age)
}
