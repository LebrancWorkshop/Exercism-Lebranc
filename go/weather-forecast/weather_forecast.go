package weather // Package weather: Using the function that related to weather. 

var CurrentCondition string // CurrentCondition: Tell The current weather condition in a city.
var CurrentLocation string // CurrentLocation: Tell The specific location in the Goblinocious. 

// func Forecast: Weather Forecasting Function. 
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition 
}
