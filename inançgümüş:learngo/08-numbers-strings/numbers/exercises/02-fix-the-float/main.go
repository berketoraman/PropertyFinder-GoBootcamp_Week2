package main

import "fmt"

// EXERCISE: Fix the Float
//  Fix the program to print 2.5 instead of 2

// EXPECTED OUTPUT
//  2.5

func main() {
	// Below solutions are correct:
	x := 5. / 2
	// x := 5 / 2.
	// x := float64(5) / 2
	// x := 5 / float64(2)

	fmt.Println(x)
}