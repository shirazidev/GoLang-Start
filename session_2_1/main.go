package main

import (
	"flag"
	"fmt"
)

func main() {
	firstName := flag.String("name", "Amir", "description")
	age := flag.Int("age", 18, "age")

	flag.Parse()
	fmt.Println("name:", *firstName, " age:", *age)

	rFlags := flag.Args()
	fmt.Printf("len of remaining flags: %d\n", len(rFlags))
	fmt.Printf("remaining flags: %s\n", rFlags)

	var name string
	fmt.Scanln(&name)
	fmt.Println("name:", name)
}
