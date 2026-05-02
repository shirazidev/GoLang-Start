package main

import "fmt"

func main() {
	s := make([]int, 3, 12)
	s = append(s, []int{1, 2, 3}...)

	fmt.Println(s)

	s = append(s, []int{3, 4, 5, 6}...)

	fmt.Println(s)
}
