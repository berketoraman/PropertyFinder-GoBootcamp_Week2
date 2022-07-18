package main

import "fmt"

func main() {
	var (
		myAge   = 30
		yourAge = 35
		average float64
	)

	//average is equal to half of the summation of myage and yourage
	average = float64(myAge+yourAge) / 2

	fmt.Println(average)
}
