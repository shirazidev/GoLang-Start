package main

import (
	"fmt"
	"math"
)

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func main() {

	var i = 21
	fmt.Println(i)

	sum := 0
	for i := 0; i <= 10; i += 2 {
		sum += i
	}
	println(sum)

	// Like while in C!
	for i < 100 {
		i += 20
	}
	println(i)

	fmt.Println("sqrts:\t", sqrt(2), sqrt(-4))

	y := 4

	/*
		when we declare some variable in a block, it is only available
		in that block and not outside of it. So here x is only available
		inside the if statement and not outside of it.
	*/

	if x := sqrt(float64(y)); x == "2" {
		fmt.Println("The square root of 4 is 2")
	}

	/*
		the first case that happens always have priority
	*/
	n := 100
	switch {
	case n > 100:
		defer fmt.Println("n > 100")
	case n < 100:
		defer fmt.Println("n < 100")
	default:
		defer fmt.Println("n == 100")
	}

}
