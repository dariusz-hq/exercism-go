// Package weather does something nice.
package weather

// CurrentCondition describes current weather condition.
var CurrentCondition string

// CurrentLocation describes current weather location.
var CurrentLocation string

// Forecast returns a string with a description for given city and weather.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
