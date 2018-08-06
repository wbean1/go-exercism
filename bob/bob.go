/*
   Package bob implements a function to address Bob.
   Bob will reply in various ways depending on how you
   address him.
*/
package bob

import (
	"regexp"
	"strings"
)

// Hey is the function for addressing Bob with a remark to which he will reply
func Hey(remark string) string {

	re := regexp.MustCompile("[?]$")
	isQuestion := re.MatchString(strings.TrimSpace(remark))

	var isYell bool
	if (strings.ToUpper(remark) == remark) && (strings.ToLower(remark) != remark) {
		isYell = true
	}

	var isNothing bool
	if strings.TrimSpace(remark) == "" {
		isNothing = true
	}

	if isQuestion && !isYell {
		return "Sure."
	}
	if isYell && !isQuestion {
		return "Whoa, chill out!"
	}
	if isYell && isQuestion {
		return "Calm down, I know what I'm doing!"
	}
	if isNothing {
		return "Fine. Be that way!"
	}
	return "Whatever."
}