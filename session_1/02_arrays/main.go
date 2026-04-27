package main

import "fmt"

func main() {
	// arrays
	var a [5]int

	a = [5]int{1, 2, 3, 4, 5}

	for i, v := range a {
		fmt.Println(i, v)
	}

	// Slice: Arrays but dynamic sized!
	var b []int

	// Built-in Functions
	b = append(b, 5)

	fmt.Println(cap(b), "\t", len(b))

	mlen := 5
	mcap := 10
	c := make([]int, mlen, mcap)
	fmt.Println(cap(c), "\t", len(c))

	f := make([]int, mlen)
	fmt.Println(cap(f), "\t", len(f))

	f = append(f, 1, 2, 3)
	fmt.Println(cap(f), "\t", len(f))

	d := []int{1, 2, 3, 4, 5}
	fmt.Println(cap(d), "\t", len(d))

	fmt.Println(d[1:4])
}
