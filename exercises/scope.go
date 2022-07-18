/*
A declaration associates a name with a program entity, such as a function or a variable. The
scope of a declaration is the part of the source code where a use of the declared name refers to that declaration.

We should not confuse scope with lifetime. The scope of a declaration is a region of the program text; it is a compile-time property. 
The lifetime of a variable is the range of time during execution when the variable can be referred to by 
other parts of the program; it is a run-time property.
*/
package main

// file scope
import "fmt"

// package scope
const ok = true

//declared again in the nested function
var declareMeAgain = 10

// package scope
func main() { // block scope starts

	var hello = "Hello"

	// hello and ok are visible here
	fmt.Println(hello, ok)
	fmt.Println("Outside function: ", declareMeAgain)
	nested()

} // block scope ends


func nested() { // block scope starts

	// declares the same variable
	// they both can exist together
	// this one only belongs to this scope
	// package's variable is still intact
	var declareMeAgain = 5
	fmt.Println("inside nested:", declareMeAgain)

} // block scope ends