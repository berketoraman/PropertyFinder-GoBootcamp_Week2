package main

import "fmt"
// EXERCISE: Iota Seasons
//  Use iota to initialize the season constants.

// HINT
//  You can change the order of the constants.

// EXPECTED OUTPUT
//  12 3 6 9

func main() {
	const (
		Spring = (iota + 1) * 3
		Summer
		Fall
		Winter
	)

	fmt.Println(Winter, Spring, Summer, Fall)
}
