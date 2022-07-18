package main

import "fmt"

func main() {
	// as you can see, we don't need to import a package
	// and we can call `hello` function here.
	//
	// this is because, package-scoped names
	// are shared in the same package
	hello()

	// but here, we can't access the fmt package without
	// importing it.
	//
	// this is because, it's in the printer.go's file scope.
	// it imports it.

	// this main func can also call bye function here
	// bye()
}

// printer.go can call this function
//
// this is because, bye function is in the package-scope
// of the main package now.
//
// main func can also call this.
func bye() {
	fmt.Println("bye bye")
}


// EXERCISE: Try the scopes
//
//  1. Create two files: main.go and printer.go
//
//  2. In printer.go:
//     1. Create a function named hello
//     2. Inside the hello function, print your name
//
//  3. In main.go:
//     1. Create the usual func main
//     2. Call your function just by using its name: hello
//     3. Create a function named bye
//     4. Inside the bye function, print "bye bye"
//
//  4. In printer.go:
//     1. Call the bye function from
//        inside the hello function

