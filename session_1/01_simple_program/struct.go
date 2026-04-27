package main

import (
	"simple/format"
	"simple/student"
)

func main() {
	// declare a variable of type student with zero values!
	//var u Student
	//u.Age = 100
	//u.StudentID = 1000
	//u.FirstName = "Mansour"
	//u.LastName = "Marzban"
	//fmt.Println(u.Age, "\n", u.StudentID, "\n", u.FirstName, u.LastName)
	//
	//var j = Student{
	//	NationalID: "229191919",
	//	Age:        22,
	//	StudentID:  108090,
	//	FirstName:  "Amir",
	//	LastName:   "Shirazi",
	//}
	//fmt.Println(j.Age, "\n", j.StudentID, "\n", j.FirstName, j.LastName)

	var p = student.Student{NationalID: "100", Age: 18, StudentID: 2098, FirstName: "Mamalizz", LastName: "Ghorbani"}
	format.Format(p)
}
