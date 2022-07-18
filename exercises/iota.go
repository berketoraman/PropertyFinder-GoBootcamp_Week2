/*
A const declaration may use the constant generator iota, which is used to create a sequence of related 
values without spelling out each one explicitly. 
In a const declaration, the value of iota begins at zero and increments by one for each item in the sequence.
*/
package main

import "fmt"

func main() {
	//This one is example of without iota.
	/*
	const (
		monday    = 0
		tuesday   = 1
		wednesday = 2
		thursday  = 3
		friday    = 4
		saturday  = 5
		sunday    = 6
	)
	fmt.Println(monday, tuesday, wednesday, thursday, friday,
		saturday, sunday)
	
	// Output: 0 1 2 3 4 5 6
	
	*/

	const (
		monday = iota 
		tuesday
		wednesday
		thursday
		friday
		saturday
		sunday
	)

	fmt.Println(monday, tuesday, wednesday, thursday, friday,
		saturday, sunday)
	// Output: 0 1 2 3 4 5 6 	
}
