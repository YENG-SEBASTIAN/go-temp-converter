package main

import (
	"fmt"
	"go-temp-converter/converter"
)

func main() {
	fmt.Println("🌡️ Temperature Converter")
	fmt.Println("------------------------")

	for {
		fmt.Print("\nEnter temperature value (or 'q' to quit): ")
		var input string
		fmt.Scanln(&input)

		if input == "q" || input == "Q" {
			fmt.Println("Exiting converter... 👋")
			break
		}

		var value float64
		_, err := fmt.Sscan(input, &value)
		if err != nil {
			fmt.Println("Invalid input! Please enter a number.")
			continue
		}

		fmt.Print("Enter current unit (C/F/K): ")
		var from string
		fmt.Scanln(&from)

		fmt.Print("Enter unit to convert to (C/F/K): ")
		var to string
		fmt.Scanln(&to)

		var result float64
		switch from + "->" + to {
		case "C->F":
			result = converter.CelsiusToFahrenheit(value)
		case "F->C":
			result = converter.FahrenheitToCelsius(value)
		case "C->K":
			result = converter.CelsiusToKelvin(value)
		case "K->C":
			result = converter.KelvinToCelsius(value)
		case "F->K":
			result = converter.FahrenheitToKelvin(value)
		case "K->F":
			result = converter.KelvinToFahrenheit(value)
		default:
			fmt.Println("Invalid conversion! Use C, F, or K.")
			continue
		}

		fmt.Println(converter.FormatResult(value, from, to, result))
	}
}
