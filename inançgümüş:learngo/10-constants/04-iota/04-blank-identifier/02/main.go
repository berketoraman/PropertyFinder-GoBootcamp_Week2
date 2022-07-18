package main

import "fmt"

func main() {
	const (
		EST = -(5 + iota) // CORRECT: -5
		MST               // CORRECT: -6
		PST               // CORRECT: -7
	)

	fmt.Println(EST, MST, PST)
	//Output: -5 -6 -7
}
