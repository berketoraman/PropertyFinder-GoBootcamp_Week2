package main

// EXERCISE: Print the Type #4
//  Print the type and value of true using fmt.Printf

// EXPECTED OUTPUT
//  Type of true is bool

import "fmt"

func main() {
	fmt.Printf("Type of %t is %[1]T\n", true)
}
