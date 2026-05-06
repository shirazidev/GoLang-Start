package main

import "fmt"

func main() {
	// pointer's zero value is nil
	var p *int
	i := 43

	fmt.Println("p", p)
	fmt.Println("i", i)

	p = &i

	fmt.Println("p after assigning i's pointer to it:", p)
	fmt.Println("value that p is refrencing to:", *p, "\n")

	d := 44
	p = &d
	fmt.Printf("&i: %p, &d: %p\n", &i, &d)
	fmt.Printf("&d: %p, p: %p, *p: %d\n", &d, p, *p)
	*p = d
	fmt.Printf("&d: %p, p: %p, *p: %d\n", &d, p, *p)
}
