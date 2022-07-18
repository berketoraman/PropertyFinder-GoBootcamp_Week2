package main
import "fmt"

// EXERCISE: Naked Expression

//  1. Try to type just "Hello" on a line.
//  2. Do not use Println
//  3. Observe the error


func main() {
	/*
		"Hello"
	*/

	// It says: "evaluted but not used"
	// Because: "Hello" literal returns a value but there isn't any statement which uses it.
	// So: You can't use expressions alone without statements.
	//The correct form is below :

	fmt.Println("Hello")
}
