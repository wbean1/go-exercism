/*
  Package reverse allows you to get the backwards
  value from your input.  Currently only strings are
  supported
*/

package reverse

// String takes a string and returns the same string reversed
func String(s string) string {
	orig := []rune(s)
	size := len(orig)
	rev := make([]rune, size)
	for i, r := range orig {
		rev[size-i-1] = r
	}
	return string(rev)
}
