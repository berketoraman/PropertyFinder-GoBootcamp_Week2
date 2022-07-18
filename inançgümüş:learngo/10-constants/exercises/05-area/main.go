package main

// EXERCISE: Area
//  Fix the following program.

// RESTRICTION
//  You should not use any variables.

// EXPECTED OUTPUT
//  area = 1250

import "fmt"

func main() {
	// Make them typeless
	const (
		width  = 25
		height = width * 2
	)

	fmt.Printf("area = %d\n", width*height)
}

