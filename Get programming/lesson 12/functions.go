package main

import "fmt"

// kelvinToCelsius converts ºK to ºC
func kelvinToCelsius(k float64) float64 {
	return k - 273.15
}

// celsiusToFahrenheit converts ºC to ºF
func celsiusToFahrenheit(c float64) float64 {
	return (c * 9.0 / 5.0) + 32.0
}

// kelvinToFahrenheit converts ºK to ºF
func kelvinToFahrenheit(k float64) float64 {
	return celsiusToFahrenheit(kelvinToCelsius(k))
}

// go run functions.go
func main() {
	kelvin := 0.0 //284.0
	celsius := kelvinToCelsius(kelvin)
	fahrenheit := celsiusToFahrenheit(celsius)
	fahrenheit2 := kelvinToFahrenheit(kelvin)

	fmt.Printf("%-.2fº K is %.2fº C\n", kelvin, celsius)
	fmt.Printf("%.2fº C is %.2fº F\n", celsius, fahrenheit)
	fmt.Printf("%.2fº K is %.2fº F\n", kelvin, fahrenheit2)
}
