package main

import (
	"fmt"
	"os"
	"unicode/utf8"
)
// EXERCISE: Count the Chars
//  1. Change the following program to work with unicode
//     characters.

// INPUT
//  "Berke"

// EXPECTED OUTPUT
//  5

func main() {
	// Currently it returns 7
	// Because, it counts the bytes...
	// It should count the runes (codepoints) instead.
	//
	// When you run it with "Berke", it should return 5 not 7.

	length := utf8.RuneCountInString(os.Args[1])
	fmt.Println(length)
}
