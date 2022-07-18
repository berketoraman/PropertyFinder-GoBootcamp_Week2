// EXERCISE: Rename imports
//
//  1- Import fmt package three times with different names
//
//  2- Print a few messages using those imports
//
// EXPECTED OUTPUT
//  hello
//  hey
//  hi

//func main() {
	// ?
	// ?
	// ?
//}

package main

import "fmt"
import f "fmt"
import fm "fmt"

func main() {
	fmt.Println("hello")
	f.Println("hey")
	fm.Println("hi")
}
