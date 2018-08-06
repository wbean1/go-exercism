/*
Package gigasecond implements the addition of
10^9 seconds (1 gigasecond) to a given time
through the use of AddGigasecond() method
*/
package gigasecond

// import path for the time package from the standard library
import "time"

// AddGigasecond should have a comment documenting it.
func AddGigasecond(t time.Time) time.Time {
	// a gigasecond is 1,000,000,000 seconds
	// so first we make a duration that represents it
	gs := (time.Duration(1000000000) * time.Second)

	// then we can simply add it to the time passed in
	// with time's Add() method
	return t.Add(gs)
}