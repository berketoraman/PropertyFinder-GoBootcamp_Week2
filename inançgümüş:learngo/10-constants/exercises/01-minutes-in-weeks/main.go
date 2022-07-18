package main

// EXERCISE: Minutes in Weeks
//  Calculate how many minutes in two weeks.

//  STEPS:
//  1. Declare `minsPerDay` constant and initialize it
//     to the number of minutes in a day

//  2. Declare `weekDays` constant and initialize it
//     to 7.

//  3. Use printf to print the total number of minutes
//     in two weeks.

// EXPECTED OUTPUT
//  There are 20160 minutes in 2 weeks.

import "fmt"

func main() {
	const (
		minsPerDay = 60 * 24
		weekDays   = 7
	)

	fmt.Printf("There are %d minutes in %d weeks.\n",
		minsPerDay*weekDays*2, 2)
}
