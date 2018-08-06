/*
Package armstrong provides a function for determining
if a number is an Armstrong number;  a number that is
equal the sum of its own digits each raised to the power
of the number of digits.

eg:

9 is; 9 = 9^1 = 9.

10 is not; 10 != 1^2 + 0^2 = 1.

153 is; 153 = 1^3 + 5^3 + 3^3 = 1 + 125 + 27 = 153.

154 is not; 154 != 1^3 + 5^3 + 4^3 = 1 + 125 + 64 = 190.
*/
package armstrong

import (
	"math"
)

// IsNumber evaluates if a provided integer is an Armstrong number
func IsNumber(i int) bool {
	digits := []int{}
	for j := i; j > 0; j = j / 10 {
		digits = append(digits, j%10)
	}
	var sum int
	for _, num := range digits {
		sum += int(math.Pow(float64(num), float64(len(digits))))
	}
	return sum == i
}
