package main

import (
	"fmt"
	time2 "time"
)

func pr(r int) int {
	fmt.Println("x: ", r)
	return r
}

func time(r int) int {
	fmt.Println("time: ", time2.Now(), r)
	return r
}

func main() {
	// function values

	//f := func(b, p int) int {
	//	return b * p
	//}
	//fmt.Println(f(2, 4))

	compute(2, pr)
	compute(6, time)

	// anonymous function
	compute(8, func(t int) int {
		fmt.Println("t: ", t)
		return t * t
	})
}

//func power(b, p int) int {
//	sum := b
//	for i := 1; i < p; i++ {
//		sum *= b
//	}
//	return sum
//}

func compute(b int, f func(r int) int) int {

	result := b * 2
	f(result)
	return result
}
