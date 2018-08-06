/*
	Package raindrops provides a function for converting
	a number into the fictional rain sounds it makes based
	on its mathematic factors.
*/
package raindrops

import (
	"fmt"
)

type rainSound struct {
	factor int
	sound  string
}

var rainSounds = [3]rainSound{
	{3, "Pling"},
	{5, "Plang"},
	{7, "Plong"},
}

// Convert returns the sounds made by a given int
// according to the factor rules defined above.
func Convert(number int) string {
	var str string
	for _, s := range rainSounds {
		if number%s.factor == 0 {
			str += s.sound
		}
	}
	if str == "" {
		return fmt.Sprint(number)
	}
	return str
}
