/*
Package leap provides functions for working with
the concept of leap years.
*/
package leap

/*
Function IsLeapYear determines if a provided year is a
leap year, or not.  A year is a leap year if it 1) is
evenly divisable by 4, AND 2) is either NOT evenly divisible
by 100 OR is evenly	divisible by 400.
*/
func IsLeapYear(year int) bool {
	return year%4 == 0 && (year%100 != 0 || year%400 == 0)
}
