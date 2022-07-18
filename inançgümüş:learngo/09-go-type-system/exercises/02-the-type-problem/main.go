package main

// EXERCISE: The Type Problem
// Solve the data type problem in the program.

// EXPECTED OUTPUT
//  width: 265 height: 265
//  are they equal? true

import "fmt"

func main() {
	var (
		width  uint16
		height uint16
	)

	width, height = 255, 265
	width += 10

	fmt.Printf("width: %d height: %d\n", width, height)
	fmt.Println("are they equal?", width == height)
}
