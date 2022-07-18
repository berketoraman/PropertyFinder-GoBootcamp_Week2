package main

import "fmt"

/*
Printf - "Print Formatter" this function allows you to format numbers, variables and strings into the first string parameter you give it
Print - "Print" This cannot format anything, it simply takes a string and print it
Println - "Print Line" same thing as Print() however it will append a newline character \n at the end.
*/
func main() {
	ops, ok, fail := 2350, 543, 433

	fmt.Println(
		"total:", ops, "- success:", ok, "/", fail,
	)

	fmt.Printf(
		"total: %d - success: %d / %d\n",
		ops, ok, fail,
	)
}
