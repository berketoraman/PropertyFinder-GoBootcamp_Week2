package main

import "fmt"

func main() {
	const min = 42

	// I've removed int from the below declaration
	// since, min's default type is int (you'll learn)
	var i = min

	var f float64 = min
	var b byte = min
	var j int32 = min
	var r rune = min

	// behind the scenes:
	// below statement equals to:
	//
	// var b byte = min
	b = byte(min)

	fmt.Println(i, f, b, j, r)
	//Output: 42 42 42 42 42
}
