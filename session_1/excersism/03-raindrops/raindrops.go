package raindrops

import "fmt"

func Convert(n int) string {
	sentence := ""
	if n%3 == 0 {
		sentence += "Pling"
	}
	if n%5 == 0 {
		sentence += "Plang"
	}
	if n%7 == 0 {
		sentence += "Plong"
	}
	if sentence == "" {
		sentence = fmt.Sprintf("%d", n)
	}
	return sentence
}
