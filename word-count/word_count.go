/*
	Package wordcount provides a function for
	counting the occurrances of words within
	strings
*/
package wordcount

import (
	"strings"
)

type Frequency map[string]int

var punctuation = []string{
	"!",
	"@",
	"#",
	"$",
	"%",
	"^",
	"&",
	"*",
	"(",
	")",
	":",
	".",
}

/*
	Function WordCount takes a phrase and returns
	a Frequency map, which maps each word contained
	in the phrase to the number of times it occurs.

	Capitalization and punctuation are (mostly) ignored.
*/
func WordCount(phrase string) Frequency {
	results := Frequency{}
	phrase = strings.ToLower(phrase)
	for _, punc := range punctuation {
		phrase = strings.Replace(phrase, punc, "", -1)
	}
	words := strings.FieldsFunc(phrase, split)
	for _, word := range words {
		word = strings.Trim(word, "'")
		results[word]++
	}
	return results
}

func split(r rune) bool {
	return r == ' ' || r == ',' || r == '\n'
}
