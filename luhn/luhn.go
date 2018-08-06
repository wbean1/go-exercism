/*
  Package luhn validates if a string is valid
  according to the luhn algorithm. See:
  https://en.wikipedia.org/wiki/Luhn_algorithm
*/
package luhn

import (
	"regexp"
	"strings"
)

// Valid determines if a string successfully passes the luhn algorithm.
func Valid(s string) bool {
	// remove spaces
	s = strings.Replace(s, " ", "", -1)

	// string must be 2 or more digits (single digits and non-digits not valid)
	re := regexp.MustCompile(`^\d\d+$`)
	if !re.MatchString(s) {
		return false
	}
	var sum int
	var nDigits = len(s)
	var parity = nDigits % 2
	for i := 0; i < nDigits; i++ {
		var digit = int(s[i] - 48)
		if i%2 == parity {
			digit *= 2
			if digit > 9 {
				digit -= 9
			}
		}
		sum += digit
	}
	return sum%10 == 0
}
