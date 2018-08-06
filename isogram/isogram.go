/*
  Package isogram supplies a function for determining
  if a string obeys the rules of an isogram; that is,
  no characters (other than - & <space>) can appear
  more than once in a word
*/

package isogram

import (
	"strings"
)

// IsIsogram tests a string to see if its an isogram
func IsIsogram(s string) bool {
	seen := make(map[string]bool)
	for _, c := range strings.ToUpper(s) {
		if seen[string(c)] {
			return false
		}
		if (string(c) != "-") && (string(c) != " ") {
			seen[string(c)] = true
		}
	}
	return true
}
