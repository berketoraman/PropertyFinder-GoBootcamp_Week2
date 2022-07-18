//Arithmetic Operators in Go
/*

+	Adds two operands	A + B gives 30
-	Subtracts second operand from the first	A - B gives -10
*	Multiplies both operands	A * B gives 200
/	Divides the numerator by the denominator.	B / A gives 2
%	Modulus operator; gives the remainder after an integer division.	B % A gives 0
++	Increment operator. It increases the integer value by one.	A++ gives 11
--	Decrement operator. It decreases the integer value by one.

*/
package main

import "fmt"

func main() {

   var a int = 120
   var b int = 12
   var c,d,e,f,g int

   c = a + b
   fmt.Println("C is equal to the summation of a and b")
   fmt.Printf("C is equal to %d\n", c )
   
   d = a - b
   fmt.Println("D is equal to the substraction of a and b")
   fmt.Printf("D is equal to %d\n", d )
   
   e = a * b
   fmt.Println("E is equal to the multiplication of a and b")
   fmt.Printf("E is equal to  %d\n", e )
   
   f = a / b
   fmt.Println("F is equal to the division of a and b")
   fmt.Printf("F is equal to %d\n", f )
   
   g = a % b
   fmt.Printf("Value of g is %d\n", g )
   
   //Increment operator. It increases the integer value by one
   a++
   fmt.Printf("Value of a is %d\n", a )
   //Decrement operator. It decreases the integer value by one
   a--
   fmt.Printf("Value of a is %d\n", a )
}