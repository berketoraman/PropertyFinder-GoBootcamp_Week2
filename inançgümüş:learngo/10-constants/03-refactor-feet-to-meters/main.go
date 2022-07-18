package main

import (
	"fmt"
	"os"
	"strconv"
	"math"
)

// EXERCISE
//
// This program uses magic numbers
// Refactor it and use named constants instead

func main() {
	// You cannot use this syntax.
	//
	// feetInMeters is declared
	//   when the declaration completes,
	//   (at the end of the whole declaration)
	//
	// And, it's not easy to read
	//
	// const feetInMeters, feetInYards float64 = 0.3048,
	//   feetInMeters / 0.9144

	const (
		feetInMeters = 0.3048
		feetInYards  = feetInMeters / 0.9144

		// cannot call runtime funcs
		// feetInYards = math.Round(feetInYards)
	)

	// Immutable: You cannot assign new values
	// feetInYards = 0.333333

	arg := os.Args[1]

	feet, _ := strconv.ParseFloat(arg, 64)

	meters := feet * feetInMeters
	yards := math.Round(feet * feetInYards)

	fmt.Printf("%g feet is %g meters.\n", feet, meters)
	fmt.Printf("%g feet is %g yards.\n", feet, yards)

	/*Output
go run main.go 205
205 feet is 62.484 meters.
205 feet is 68 yards.
*/
}