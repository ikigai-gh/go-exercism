// Package weather contains information about the weather forecast.
package weather

// CurrentCondition is current weather condition.
var CurrentCondition string

// CurrentLocation is current location.
var CurrentLocation string

// Forecast. This function returns weather forecast based on provided city and condition.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
