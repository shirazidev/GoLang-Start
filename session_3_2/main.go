package main

import (
	"fmt"
	"session_3_2/student"
)

func main() {
	s := student.Student{Name: "Amir"}
	s2 := student.Student{Name: "Amir"}
	fmt.Println(s.Id())
	fmt.Println(s2.Id())
}
