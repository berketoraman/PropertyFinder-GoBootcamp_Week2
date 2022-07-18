package main

import "fmt"

// EXERCISE: Convert the Types
//  Convert the variables to appropriate types.

// EXPECTED OUTPUT
//  325.5 299.5

func main() {
	// DONT TOUCH THIS:
	type (
		Temperature float64
		Celsius     Temperature
		Fahrenheit  Temperature
	)

	// DONT TOUCH THIS:
	var (
		celsius       Celsius     = 15.5
		fahr          Fahrenheit  = 59.9
		celsiusDegree Temperature = 10.5
		fahrDegree    Temperature = 2.5
		factor                    = 2.
	)

    celsius *= Celsius(float64(celsiusDegree) * factor)
	fahr *= Fahrenheit(float64(fahrDegree) * factor)

	fmt.Println(celsius, fahr)
}


