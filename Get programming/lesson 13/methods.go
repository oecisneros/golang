package main

import "fmt"

type celsius float64
type fahrenheit float64
type kelvin float64

// toCelsius converts ºK to ºC
func (k kelvin) toCelsius() celsius {
	return celsius(k - 273.15)
}

// toFahrenheit converts ºK to ºF
func (k kelvin) toFahrenheit() fahrenheit {
	return k.toCelsius().toFahrenheit()
}

// toKelvin converts ºC to ºK
func (c celsius) toKelvin() kelvin {
	return kelvin(c + 273.15)
}

// toFahrenheit converts ºC to ºF
func (c celsius) toFahrenheit() fahrenheit {
	return fahrenheit((c * 9.0 / 5.0) + 32.0)
}

// toCelsius converts ºF to ºC
func (f fahrenheit) toCelsius() celsius {
	return celsius((f - 32) / 1.8)
}

// toKelvin converts ºF to ºK
func (f fahrenheit) toKelvin() kelvin {
	return f.toCelsius().toKelvin()
}

func main() {
	var k1 kelvin = 0.0 //284.0
	c1 := k1.toCelsius()
	f1 := c1.toFahrenheit()
	f2 := k1.toFahrenheit()
	k2 := f2.toKelvin()

	fmt.Printf("%8.2fº K is %8.2fº C\n", k1, c1)
	fmt.Printf("%8.2fº C is %8.2fº F\n", c1, f1)
	fmt.Printf("%8.2fº K is %8.2fº F\n", k1, f2)
	fmt.Printf("%8.2fº F is %8.2fº K\n", f2, k2)
}
