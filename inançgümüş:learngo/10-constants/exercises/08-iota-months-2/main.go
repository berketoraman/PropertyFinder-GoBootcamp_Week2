package main

import "fmt"

// EXERCISE: Iota Months #2
//  1. Initialize multiple constants using iota.
//  2. Please follow the instructions inside the code.

// EXPECTED OUTPUT
//  1 2 3

func main() {
	const (
		_   = iota // 0
		Jan        // 1
		Feb        // 2
		Mar        // 3
	)

	fmt.Println(Jan, Feb, Mar)
}
