package main

import (
	"fmt"
	"os"
	"strings"
	"unicode/utf8"
)
// EXERCISE: Improved Banger
//  Change the Banger program the work with Unicode
//  characters.

// INPUT
//  "İNANÇ"

// EXPECTED OUTPUT
//  Berke !!!!!

func main() {
	msg := os.Args[1]
	l := utf8.RuneCountInString(msg)

	s := msg + strings.Repeat("!", l)

	fmt.Println(s)
}
