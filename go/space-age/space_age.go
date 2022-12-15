/*
Package space recalculates times in seconds to years on planets of our solar system
 */
package space

const MercuryOP = 0.2408467
const VenusOP = 0.61519726
const EarthOP = 1.0
const MarsOP = 1.8808158
const JupiterOP = 11.862615
const SaturnOP = 29.447498
const UranusOP = 84.016846
const NeptuneOP = 164.79132
const EarthOPSeconds =  31557600

var planetOrbitalPeriods = map[Planet]float64{"Mercury" : MercuryOP, "Venus" : VenusOP, "Earth" : EarthOP, "Mars" : MarsOP, "Jupiter" : JupiterOP, "Saturn" : SaturnOP, "Uranus" : UranusOP, "Neptune" : NeptuneOP}

type Planet string

// Returns years from seconds on particular planet
func Age(seconds float64, planet Planet) float64 {
	return seconds / planetOrbitalPeriods[planet] / EarthOPSeconds
}
