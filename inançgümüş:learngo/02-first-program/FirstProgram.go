//File write-your-first-program from İnanç Gümüş's learngo repository.

//İf ve move the package and import statement to the bottom of the file the program does not work because it needs package declaration at the beginning.
package main
import "fmt"

func main() {

	fmt.Println("Hi! I want to be a Gopher!")
	fmt.Println("This is Berke Toraman")

	//Used Print instead of Println
	fmt.Print("Hi! I want to be a Gopher!")
	fmt.Print("This is Berke Toraman\n")
	//Print does not add any new line while Println adds a new line after displaying the text.

	//fmt.Println(This is Berke Toraman)
	//if we remove the double quotes from a string literal the program does not work.It doesn't pass the message to Println.

	//Call Println or Print with multiple values by separating them using commas
	fmt.Println("selam","berke","toraman",) //Output: selam berke toraman
	fmt.Print("berke","toraman","selam") //Output: berketoramanselam	
}

/////////////////////////////////////////////////////// Questions and Answers ///////////////////////////////////////////////////////

//Where should you save your Go source code?
//Answer: Under $GOPATH/src 

//Do you need to set $GOPATH?
//Answer: No, it's stored under my user path 

//How can you print $GOPATH?
//Answer: Using go env GOPATH command 

//Which keyword below that you need use to define a package?
//Answer: package

//What is the purpose of using package main in the following program?
//Answer: To create an executable Go program 

//What is the purpose of func main in the following program?
//Answer: It allows Go to start executing the program

//What is the purpose of import "fmt" in the following program?
//Answer: It imports the fmt package; so you can use its functionalities

//Which keyword is used to declare a new function?
//Answer: func

//What is a function?
//Answer: It's like a mini program. It's a reusable and executable block of code

//Do you have to call the main function yourself?
//Answer: No, Go calls the main function automatically.

//Do you have to call a function to execute it?
//Answer: Yes, so that, Go can execute that function. (except the main func)

/*
What does the following program print?

package main
func main() {
}

Answer: It's a correct program but it doesn't print anything 
*/


/*
What does this program print?

package main
func main() {
    fmt.Println(Hi! I want to be a Gopher!)
}

Answer: This program is incorrect 
It doesn't pass the message to Println wrapped between double-quotes. It should be like: fmt.Println("Hi! I want to be a Gopher")
It doesn't import "fmt" package.
*/


/*
What does this program print?

package main
import "fmt"

func main() {
    fmt.Println("Hi there!")
}

Answer: Hi there! 
*/


//What is the difference between go build and go run?
//Answer: Go run both compiles and runs a program; whereas go build just compiles it.

//Go saves the compiled code in a directory. What is the name of that directory?
//Answer: The same directory where you call go build 

//Which is true for runtime?
//Answer: It happens when your program starts running on a computer

//Which is true for the compile-time?
//Answer: It happens while your program is being compiled 

//When can a Go program print a message to the console?
//Answer: While it runs (after compile-time). 

