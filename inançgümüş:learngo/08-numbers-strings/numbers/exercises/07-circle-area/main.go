package main

// EXERCISE: Circle Area
//  Calculate the area of a circle from the given radius

// CIRCLE AREA FORMULA
//  area = πr²
//  https://en.wikipedia.org/wiki/Area#Circles

// HINT
//  For PI you can use `math.Pi`

// EXPECTED OUTPUT
//  radius: 10 -> area: 314.1592653589793

// BONUS EXERCISE!
//  1. Print the area as 314.16
//  2. To do that you need to use the correct Printf verb :)
//      Instead of `%g` verb below.

//    EXPECTED OUTPUT
//     radius: 10 -> area: 314.16

import (
	"fmt"
	"math"
)

func main() {
	var (
		radius = 10.
		area   float64
	)

	area = math.Pi * radius * radius

	fmt.Printf("radius: %g -> area: %.2f\n",
		radius, area)

	// ALTERNATIVE:
	// math.Pow calculates the power of a float number
	// area = math.Pi * math.Pow(radius, 2)
}