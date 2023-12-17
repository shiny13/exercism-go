// Package weather provides tools to forecast the weather.
package weather

// CurrentCondition represents the current condition of the weather.
var CurrentCondition string

// CurrentLocation shows the current location of the weather.
var CurrentLocation string

// Forecast returns the current location and temperature.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
