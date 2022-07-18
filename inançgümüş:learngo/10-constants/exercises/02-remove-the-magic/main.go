package main

import "fmt"

// EXERCISE: Remove the Magic
//  Get rid of the magic numbers in the following code.

// RESTRICTIONS
//  1. You should declare 3 constants named:
//       hoursInDay, daysInWeek, and hoursPerWeek

//  2. And, hoursPerWeek constant should be initialized
//     using hoursInDay and daysInWeek constants.

// EXPECTED OUTPUT
//  There are 840 hours in 5 weeks.


func main() {
	const (
		hoursInDay   = 24
		daysInWeek   = 7
		hoursPerWeek = hoursInDay * daysInWeek
	)

	fmt.Printf("There are %d hours in %d weeks.\n",
		hoursPerWeek*5, 5)
}
