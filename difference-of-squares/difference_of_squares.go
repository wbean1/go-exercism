/*
  Package diffsquares provide functions for determining
  the SquareOfSums, SumOfSquares, and the Difference between
  the two, operating on the first n natural integers
*/
package diffsquares

// SquareOfSums computes (1 + 2 + ... + n)^2
func SquareOfSums(n int) int {
	sum := 0
	for i := n; i > 0; i-- {
		sum += i
	}
	return sum * sum
}

// SumOfSquares computes 1^2 + 2^2 + ... + n^2
func SumOfSquares(n int) int {
	sum := 0
	for i := n; i > 0; i-- {
		sum += i * i
	}
	return sum
}

// Difference computes SquareOfSums(n) - SumOfSquares(n)
func Difference(n int) int {
	return SquareOfSums(n) - SumOfSquares(n)
}
