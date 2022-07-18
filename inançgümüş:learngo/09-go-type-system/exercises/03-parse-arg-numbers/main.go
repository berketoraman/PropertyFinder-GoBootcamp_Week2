package main

import (
	"fmt"
	"os"
	"strconv"
)
// EXERCISE: Parse Arg Numbers
//  Use strconv.ParseInt function to get int8, int16, and
//  int32, and int64 values from command-line.

// HINT
//  The third argument to ParseInt function represents
//  the bitsize.
//  So, giving it 8 returns an int8 convertable value;
//  whereas 16 returns an int16 convertable value.
//
//  Please explore the documentation of ParseInt function
//  and learn how it works.
//
// EXPECTED OUTPUT
//   When runned like this:
//     go run main.go 50 25000 2000000 50000000000000000 00000100
//
//   It should return this:
//     int8 value is : 50
//     int16 value is: 25000
//     int32 value is: 2000000
//     int64 value is: 50000000000000000
//     00000100 is: 4
// ---------------------------------------------------------

func main() {
	
	val, _ := strconv.ParseInt(os.Args[1], 10, 8)
    fmt.Println("int8 value is :", int8(val))

	
	// 2nd argument is an int16
	val, _ = strconv.ParseInt(os.Args[2], 10, 16)
	fmt.Println("int16 value is:", int16(val))

	// 3rd argument is an int32
	val, _ = strconv.ParseInt(os.Args[3], 10, 32)
	fmt.Println("int32 value is:", int32(val))
	
	// 4th argument is an int64
	// Remember ParseInt returns an int64
	val, _ = strconv.ParseInt(os.Args[4], 10, 64)
	fmt.Println("int64 value is:", val)
	
	// 5th argument is a number in bits. And its int8.
	// ParseInt(.., 2, ...) -> 2 means binary base
	val, _ = strconv.ParseInt(os.Args[5], 2, 8)
	fmt.Printf("%s is: %d\n", os.Args[5], int8(val))
}

