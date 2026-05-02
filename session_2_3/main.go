package main

import "fmt"

func main() {

	// variadic functions
	prt("h")
	prt("salam", "man", "amiram")

}

func prt(values ...string) {
	pr(values...)
}

func pr(values ...string) {
	for _, value := range values {
		fmt.Println(value)
	}
}
