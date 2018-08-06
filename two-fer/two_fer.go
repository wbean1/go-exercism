/*
	Package twofer provides a function for describing
	how two things will be shared between two people.
*/
package twofer

import "fmt"

/*
	ShareWith takes a name of a person with whom to
	share, and returns a string decribing how two things
	will be divvied up.  One person is assumed to be "me".
	If an empty name is provided, the other person is
	assumed to be "you".
*/
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}
	return fmt.Sprintf("One for %s, one for me.", name)
}
