// Package weather provides information about the weather conditions of a particular place or location per time.
package weather

var (
    // CurrentCondition holds the value of the weather condition at a particular point in time.
	CurrentCondition string
    // CurrentLocation specifies the location whose weather condition is being addressed.
	CurrentLocation  string
)

// Forecast takes in parameters city and condition, then returns information about the current location and the prevailing weather.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
