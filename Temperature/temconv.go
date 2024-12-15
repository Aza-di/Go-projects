package main

import (
	"fmt"
)

// convert celsius to fahrenheit: (celsius * 9/5) + 32
// convert fahrenheit to celsius: (fahrenheit - 32) * 5/9
func celsiusToFahrenheit(celsius float64) float64 {
	return (celsius * 9 / 5) + 32

}

func fahrenheitToCelsius(fahrenheit float64) float64 {
	return (fahrenheit - 32) * 5 / 9
}

func main() {
	fmt.Println("Temperature converter")
	fmt.Println("Celsius to Fahrenheit")
	fmt.Println("Fahrenheit to Celsius")
	fmt.Print("Enter the conversion type (1 or 2): ")

	var choice int
	_, err := fmt.Scanln(&choice)
	if err != nil {
		fmt.Println("Invalid input")
		return
	}
	if choice != 1 && choice != 2 {
		fmt.Println("Provide either 1 or 2")
		return
	}

	var temp float64
	fmt.Print("Enter temperature value: ")
	_, err = fmt.Scanln(&temp)
	if err != nil {
		fmt.Println("Invalid input")
		return
	}

	switch choice {
	case 1:
		result := celsiusToFahrenheit(temp)
		fmt.Printf("%.2f degrees celsius = %.2f fahrenheit\n", temp, result)
	case 2:
		result := fahrenheitToCelsius(temp)
		fmt.Printf("%.2f farenheit = %.2f celsius degrees\n", temp, result)

	}
}
