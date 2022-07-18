package main

import "fmt"

func main() {
	const (
		monday = iota + 1
		tuesday
		wednesday
		thursday
		friday
		saturday
		sunday
	)

	fmt.Println(monday, tuesday, wednesday, thursday, friday,
		saturday, sunday)
	// Output: 1 2 3 4 5 6 7 	
}
