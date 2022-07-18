package main

// EXERCISE: TAU
//  Fix the following program and print the TAU number.

// HINT
//  You can use %g verb for printing tau.

// EXPECTED OUTPUT
//  tau = 6.283185307179586

import "fmt"

func main() {
	const (
		pi  = 3.14159265358979323846264
		tau = pi * 2
	)

	fmt.Printf("tau = %g\n", tau)
}
