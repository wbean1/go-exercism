/*
	Package accumulate provides a mechanism for performing an arbitrary
	input function (called an accumulator) on an input set
*/

package accumulate

// Function Accumulate performs a provided function on a set of input
// strings, and returns the resulting strings.  The function provided
// must have a signature `func(string) string`
func Accumulate(inputs []string, function func(string) string) []string {
	results := make([]string, len(inputs))
	for i, x := range inputs {
		results[i] = function(x)
	}
	return results
}
