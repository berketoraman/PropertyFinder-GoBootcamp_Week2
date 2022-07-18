package main

// EXERCISE: Incdecs

//  1. Increase the `counter` 5 times
//  2. Decrease the `factor` 2 times
//  3. Print the product of counter and factor

// RESTRICTION
//  Use only the incdec statements

// EXPECTED OUTPUT
//  -75

func main() {
	counter, factor := 45, 0.5

	counter++
	counter++
	counter++
	counter++
	counter++

	factor--
	factor--

	fmt.Println(float64(counter) * factor)
}

