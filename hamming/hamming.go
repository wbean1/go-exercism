/*
  Package hamming provides a function for computing
  the hamming distance between two dna sequences
*/

package hamming

import (
	"errors"
)

// Distance calculates the hamming distance between two dna
// sequences represented as strings.  The sequences must be
// of indentical length.
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return -1, errors.New("hamming: strings must be same length")
	}
	d := 0
	for i := range a {
		if a[i] != b[i] {
			d++
		}
	}
	return d, nil

}
