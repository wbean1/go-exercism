/*
  Package letter contains functions for computing
  the frequency of all characters in a given string
*/
package letter

type FreqMap map[rune]int

var doneMaps = make(chan FreqMap)

func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

// ConcurrentFrequency takes a list of strings and,
// using goroutines for concurrency, computes a FreqMap
// for each string, then returns a totaled FreqMap
func ConcurrentFrequency(s []string) FreqMap {
	for _, str := range s {
		go func(str string) {
			doneMaps <- Frequency(str)
		}(str)
	}

	// I tried like 10 different ways to also process
	// the results concurrently, but it required putting
	// a mutex around the totalled map otherwise I'd get:
	// `fatal error: concurrent map read and map write`
	// Using mutexes made it slower than just sequential
	// processing, so... shrug
	crushmap := FreqMap{}
	for i := 0; i < len(s); i++ {
		x := <-doneMaps
		for r, i := range x {
			crushmap[r] += i
		}
	}
	return crushmap
}
