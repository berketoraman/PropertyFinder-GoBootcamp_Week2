package main

// EXERCISE: Print Your Fullname
//  1. Get your name and lastname from the command-line
//  2. Print them using Printf

// EXAMPLE INPUT
//  Inanc Gumus

// EXPECTED OUTPUT
//  Your name is Inanc and your lastname is Gumus.

import (
	"fmt"
)

func main() {
	// WARNING: This program will error
	//if you don't pass your name and lastname

	name, lastname := "Berke", "Toraman"
	msg := "Your name is %s and your lastname is %s.\n"

	fmt.Printf(msg, name, lastname)

}
