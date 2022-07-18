package main
// EXERCISE: Windows Path
//  1. Change the following program
//  2. It should use a raw string literal instead

// HINT
//  Run this program first to see its output.
//  Then you can easily understand what it does.

// EXPECTED OUTPUT
//  Your solution should output the same as this program.
//  Only that it should use a raw string literal instead.

import "fmt"

func main() {
	// this one uses a raw string literal
	path := `c:\program files\duper super\fun.txt
c:\program files\really\funny.png`

	fmt.Println(path)
}
