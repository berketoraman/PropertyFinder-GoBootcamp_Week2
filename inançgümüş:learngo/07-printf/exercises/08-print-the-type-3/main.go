package main

// EXERCISE: Print the Type #3
//  Print the type and value of "hello" using fmt.Printf

// EXPECTED OUTPUT
// 	Type of hello is string

import "fmt"

func main() {
	fmt.Printf("Type of %s is %[1]T\n", "hello")
}