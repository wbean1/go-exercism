/*
  Package space provides a function for calculating
  the Age of a person on a given Planet using its
  relative orbital cycle.  In Mercury years, you are
  much older (and maybe wiser)
*/

package space

type Planet = string

var secondsInEarthYear = 31557600.0
var relativeOrbitalPeriod = map[Planet]float64{
	"Earth":   1.0,
	"Mercury": 0.2408467,
	"Venus":   0.61519726,
	"Mars":    1.8808158,
	"Jupiter": 11.862615,
	"Saturn":  29.447498,
	"Uranus":  84.016846,
	"Neptune": 164.79132,
}

// function Age takes a number of seconds and a Planet, and returns
//   the corresponding number of orbital years
func Age(s float64, p Planet) float64 {
	return (s / (secondsInEarthYear * relativeOrbitalPeriod[p]))
}
