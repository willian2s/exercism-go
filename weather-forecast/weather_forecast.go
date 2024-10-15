// Package weather represents ...
package weather

// CurrentCondition ...
var CurrentCondition string

// CurrentLocation ...
var CurrentLocation string

// Forecast returns a string with the given city and condition.
// It also sets the CurrentLocation and CurrentCondition variables.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
