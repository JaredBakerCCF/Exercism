// Exercism GoLang commenting practice

// Package weather provides the source files to the tools required to display the weather
// at the specified location.
package weather

// CurrentCondition represents a real-time status of the weather.
var CurrentCondition string

// CurrentLocation represents a location based off of user's current geographical location.
var CurrentLocation string

// Forecast function takes in 2 parameters (city and condition)
// and displays the current statuses for both.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
