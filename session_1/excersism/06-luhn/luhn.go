package luhn

import "unicode"

func Valid(number string) bool {
	sum := 0
	double := false
	digits := 0

	for i := len(number) - 1; i >= 0; i-- {
		r := rune(number[i])
		if r == ' ' {
			continue
		}
		if !unicode.IsDigit(r) {
			return false
		}
		n := int(r - '0')
		digits++
		if double {
			n *= 2
			if n > 9 {
				n -= 9
			}
		}
		sum += n
		double = !double
	}
	if digits <= 1 {
		return false
	}
	return sum%10 == 0
}
