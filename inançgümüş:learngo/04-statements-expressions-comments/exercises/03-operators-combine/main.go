package main

import "fmt"

// EXERCISE: Operators combine the expressions

//  Print the expected output below using the string concatenation operator.
// HINT
//  Use + operator multiple times to create "Hello!!!?".

// EXPECTED OUTPUT
//  "Hello!!!?"

func main() {
    // Operators combine multiple expressions together
	// as if there's a single expression.
	fmt.Println("Hello!" + "!" + "!" + "?")
}
