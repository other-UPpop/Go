// Package weather provides weather forecasting functionalities.
package weather

// CurrentCondition holds the current weather condition.
var CurrentCondition string

// CurrentLocation holds the current location for which the weather is being reported.
var CurrentLocation string

// Forecast returns a string describing the current weather condition in a given city.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
