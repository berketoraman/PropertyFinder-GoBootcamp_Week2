package main

// file scope
import "fmt"

// package scope
const ok = true

// package scope
func main() { // block scope starts

	var hello = "Hello"

	// hello and ok are visible here
	fmt.Println(hello, ok)

} // block scope ends

////////////////////// QUESTION AND ANSWERS ////////////////////////////

/*

What's a scope?
Answer: The visibility of the declared names

Which name below is package scoped?
Answer: enabled

Which name below is file scoped?
Answer: fmt 

Which name below is in the scope of the block() func?
Answer: counter 

Can block() see enabled name?
Answer: Yes: It's in the package scope

Can other files in awesome package see counter name?
Answer: No: It's in the block scope of block() 

Can other files in awesome package see fmt name?
Answer: No: It's in the file scope CORRECT

What happens if you declare the same name in the same scope as this:

package awesome

import "fmt"

// declared twice in the package scope
var enabled bool
var enabled bool

func block() {
    var counter int
    fmt.Println(counter)
}

I can't do that. It's already been declared at the package scope.

What happens if you declare the same name in another scope like this:

package awesome

import "fmt"

// declared at the package scope
var enabled bool

func block() {
    // also declared in the block scope
    var enabled bool

    var counter int
    fmt.Println(counter)
}
The newly declared name will override the previous one.
 
*/