/*
  Package grains provides functions to compute the number
  of grains on a certain square of a chess board, given that
  each square has 2^n grains (where n is the square index 0-63).
*/
package grains

import (
	"errors"
)

// Square computes the number of grains on a given square (indexed from 1-64).
func Square(i int) (uint64, error) {
	if (i < 1) || (i > 64) {
		return 0, errors.New("grains: invalid input to Square(); must be 1-64 inclusive")
	}
	return 1 << uint64(i-1), nil
}

// Total computes the total number of grains on the entire chessboard.
func Total() uint64 {
	var sum uint64 = 1
	for i := 64; i > 1; i-- {
		sum += 1 << uint64(i-1)
	}
	return sum
}
