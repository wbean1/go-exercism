/*
	Package piglatin encodes the rules for converting
	a phrase into its piglatin equivalents.  The package
	provides a one-way function for converting an english
	phrase to pig latin.
*/
package piglatin

import (
	"fmt"
	"regexp"
	"strings"
)

// Sentence takes a phrase (single word, or space separated
// words) and returns the similar phrase converted to
// pig latin accoding to rules encoded in the package.
func Sentence(phrase string) string {
	words := strings.Split(phrase, " ")
	var latins = make([]string, len(words))
	for i, word := range words {
		latins[i] = piggify(word)
	}
	return fmt.Sprint(strings.Join(latins, " "))
}

// piggify is a helper function which operates on a single
// word.
func piggify(word string) string {
	// rule #1, if begins with vowel, just add 'ay'
	re := regexp.MustCompile("^(a|e|i|o|u|xr|yr|yt)")
	if re.MatchString(word) {
		word += "ay"
		return word
	}

	// rule #3, handling qu squ
	re = regexp.MustCompile("^(.{0,1}qu)(.*)")
	matches := re.FindStringSubmatch(word)
	if matches != nil {
		return matches[2] + matches[1] + "ay"
	}

	// rule #4, y's are no fun
	re = regexp.MustCompile("^([^aeiou]+)(y.*)")
	matches = re.FindStringSubmatch(word)
	if matches != nil {
		return matches[2] + matches[1] + "ay"
	}

	// rule #2, move letters before vowel to end, then add 'ay'
	re = regexp.MustCompile("^([^aeiou]*)(.*)")
	matches = re.FindStringSubmatch(word)
	if matches != nil {
		return matches[2] + matches[1] + "ay"
	}

	return word
}
