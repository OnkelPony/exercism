// Package weather provides tools to display weather conditions in the city in nice manner.
package weather

// CurrentCondition represents weather condition, which is now.
var CurrentCondition string

// CurrentLocation represents place on earth.
var CurrentLocation string

// Forecast takes city and condition as a parameter and returns string representation of weather conditions
// at the given place.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
