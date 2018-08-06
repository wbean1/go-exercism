/*
	Package summultiples provides a function for
	determining the sum of multiples given certain
	constraints.
*/
package summultiples

/*
	Function SumMultiples takes a limit number, and
	a set of divisor numbers, and computes the sum
	of all numbers less than the limit which are evenly
	divided by at least one of the divisor numbers.

	For example, SumMultiples(20, 3, 5) would find
	`3 + 5 + 6 + 9 + 10 + 12 + 15 + 18 = 78`
*/
func SumMultiples(limit int, divisors ...int) int {
	multiples := make(map[int]int)
	for _, i := range divisors {
		for j := i; j < limit; j += i {
			multiples[j] = j
		}
	}
	var sum int
	for _, mult := range multiples {
		sum += mult
	}
	return sum
}
