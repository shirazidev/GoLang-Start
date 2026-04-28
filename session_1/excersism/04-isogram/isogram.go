package isogram

import (
	"strings"
	"unicode"
)

func IsIsogram(s string) bool {
	word := strings.ToLower(s)
	ite := make(map[rune]bool)
	for _, w := range word {
		if w == ' ' || w == '-' {
			continue
		}
		if unicode.IsLetter(w) {
			if ite[w] {
				return false
			}
			ite[w] = true
		}
	}
	return true
}
