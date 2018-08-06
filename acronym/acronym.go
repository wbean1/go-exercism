/*
  Package acronym provides a function for creating
  an all-capital Abbreviation of an input string
*/
package acronym

import (
	"strings"
)

// Abbreviate creates an acronym of the input string
func Abbreviate(s string) string {

	words := strings.FieldsFunc(s, Split)
	str := make([]string, len(words))
	for i := range words {
		str[i] = strings.ToUpper(string(words[i][0]))
	}
	return strings.Join(str, "")
}

// Split provides the rules for what characters are used to split words
// in Abbreviate's input to form the Abbreviation
func Split(r rune) bool {
	return r == ' ' || r == '-'
}