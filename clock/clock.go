/*
	Package clock implements a simple 24 hour clock
	that can add & subtract minutes, be represented
	by a string, and be directly compared.
*/
package clock

import (
	"fmt"
)

type Clock struct {
	Hour   int
	Minute int
}

// New returns a new Clock.  Input validation occurs here.
func New(hour, minute int) Clock {

	// this adds or subtracts hours if >60 <-60 min
	hour += minute / 60
	minute = minute % 60

	// if we're left with negative minutes, take off an hour
	if minute < 0 {
		hour--
		minute += 60
	}

	// we don't care about days, so just find the remainder for hours
	hour = hour % 24

	// if we're left with negative hours, just add 24
	if hour < 0 {
		hour += 24
	}

	c := Clock{Hour: hour, Minute: minute}
	return c
}

// String returns a string representation of the clock in hh:mm format.
func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.Hour, c.Minute)
}

// Add adds some number of minutes to a clock, and returns the new clock.
func (c Clock) Add(minutes int) Clock {
	return New(c.Hour, c.Minute+minutes)
}

// Subtract subtracts some number of minutes to a clock, and returns the new clock.
func (c Clock) Subtract(minutes int) Clock {
	return New(c.Hour, c.Minute-minutes)
}
