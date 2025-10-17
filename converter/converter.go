package converter

import "fmt"

// CelsiusToFahrenheit converts Celsius to Fahrenheit
func CelsiusToFahrenheit(c float64) float64 {
	return c*9/5 + 32
}

// FahrenheitToCelsius converts Fahrenheit to Celsius
func FahrenheitToCelsius(f float64) float64 {
	return (f - 32) * 5 / 9
}

// CelsiusToKelvin converts Celsius to Kelvin
func CelsiusToKelvin(c float64) float64 {
	return c + 273.15
}

// KelvinToCelsius converts Kelvin to Celsius
func KelvinToCelsius(k float64) float64 {
	return k - 273.15
}

// FahrenheitToKelvin converts Fahrenheit to Kelvin
func FahrenheitToKelvin(f float64) float64 {
	c := FahrenheitToCelsius(f)
	return CelsiusToKelvin(c)
}

// KelvinToFahrenheit converts Kelvin to Fahrenheit
func KelvinToFahrenheit(k float64) float64 {
	c := KelvinToCelsius(k)
	return CelsiusToFahrenheit(c)
}

// FormatResult nicely prints the conversion
func FormatResult(input float64, from, to string, result float64) string {
	return fmt.Sprintf("%.2f %s = %.2f %s", input, from, result, to)
}
