package main

import "fmt"

func main() {
	fmt.Println("Hello!")

/* You can access functions from other files
which are in the same package
*/
// Here, `main()` can access `bye()` and `hey()`

/* It's because bye.go, hey.go and main.go
are in the main package.
*/
bye()
hey()
}

///////////////////////////////////////	QUESTION AND ANSWERS ////////////////////////////////////
/*
Where to store the source code files that belong to a package?
Answer: In a single directory 

Why a package clause is used in a Go source-code file?
Answer: It's used for letting Go know that the file belongs to a package

Where you should put the package clause in a Go source-code file?
AAnswer:s the first code in a Go source code file 

How many times you can use package clause for a single source code file?
Answer: Once 

Which one is a correct usage of package clause?
Answer: package main


How to run multiple Go files?
Answer: go run *.go 

Which one below is a correct package type in Go?
Answer: Executable package 

Which package type go run can execute?
Answer: Executable package

Which package type that go build can compile?
Answer: Both of executable and library packages 

Which one is an executable package?
Answer: package main with func main

Which one is a library package?
Answer: package lib

Which package is used for an executable Go program?
Answer: Executable package 

Which package is used for reusability and can be imported?
Answer: Library package 

*/