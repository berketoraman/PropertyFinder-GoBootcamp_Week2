package main

// EXERCISE: Raw Concat
//  1. Initialize the name variable
//     by getting input from the command line
//  2. Use concatenation operator to concatenate it
//     with the raw string literal below

// NOTE
//  You should concatenate the name variable in the correct
//  place.

// EXPECTED OUTPUT
//  Let's say that you run the program like this:
//   go run main.go inanç

//  Then it should output this:
//   hi Berke!
//   how are you?

import (
	"fmt"
	"os"
)
func main() {
	name := os.Args[1]	
		msg := `hi ` + name + `!
	how are you?`
	
	fmt.Println(msg)
}

//go run main.go "Berke"
	

