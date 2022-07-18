package main

import (
	//"fmt"
	"os"
	"strings"
)

// NOTE: You should always pass it at least one argument

func main() {
	msg := os.Args[1]
	// it's important to calculate things only once
	// so, do not call the repeat function twice
	// calling it once is enough
	marks := strings.Repeat("!", len(msg))
	s := marks + msg + marks
	s = strings.ToUpper(s)
	
	// you can also type this program more concisely
	// like this:
	//l := len(msg)
	//s := msg + strings.Repeat("!", l)
	//s = strings.ToUpper(s)
	//fmt.Println(s)
}
