/*
	Package proverb constructs a personalized
	message for your proverb consuming enjoyment.
*/
package proverb

import "fmt"

// Proverb returns lines of a proverb constructed from a
// list of input words.
func Proverb(rhyme []string) []string {
	var ret = make([]string, len(rhyme))
	for i := range rhyme {
		if i == len(rhyme)-1 {
			ret[i] = fmt.Sprintf("And all for the want of a %s.", rhyme[0])
		} else {
			ret[i] = fmt.Sprintf("For want of a %s the %s was lost.", rhyme[i], rhyme[i+1])
		}
	}
	return ret
}
