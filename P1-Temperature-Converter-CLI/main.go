package main

import (
	"fmt"
	converter "temperature-converter-cli/converter"
)

func main() {
	fmt.Println("Temperature Converter")
	fmt.Println("---------------------")
	fmt.Println("Select conversion type:")
	fmt.Println("1. Celsius to Fahrenheit")
	fmt.Println("2. Fahrenheit to Celsius")
	fmt.Print("Enter your choice: (1 or 2): ")

	var choice int
	fmt.Scanln(&choice)

	var Temperature float64
	fmt.Print("Enter the temperature value to convert: ")
	fmt.Scanln(&Temperature)

	var converted float64
	var label string
	switch choice {
	case 1:
		converted = converter.CelsiusToFahrenheit(Temperature)
		label = "Fahrenheit"
	case 2:
		converted = converter.FahrenheitToCelsius(Temperature)
		label = "Celsius"
	default:
		fmt.Println("Invalid choice")
		return
	}

	fmt.Printf("Result: %.2f degrees %s\n", converted, label)
}
