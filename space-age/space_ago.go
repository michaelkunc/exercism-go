// package space age returns age in years for a given planet

package space

type Planet = string

var earthSeconds = 31557600.0
var orbitalPeriods = map[Planet]float64{
	"Earth":   1.0,
	"Mercury": 0.2408467,
	"Venus":   0.61519726,
	"Mars":    1.8808158,
	"Jupiter": 11.862615,
	"Saturn":  29.447498,
	"Uranus":  84.016846,
	"Neptune": 164.79132,
}

// Age takes seconds and a planet and returns age in years
func Age(seconds float64, planet string) float64 {
	onPlanetSeconds := earthSeconds * orbitalPeriods[planet]
	return seconds / onPlanetSeconds
}
