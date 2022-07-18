/*
Type conversion happens when we assign the value of one data type to another.
There is no typecasting word or terminology in Golang. If you will try to search Type Casting in Golang Specifications or Documentation, you will find nothing.
There is only Type Conversion. In Other programming languages, typecasting is also termed as the type conversion.
What is the need for Type Conversion? 
Well, if you need to take advantage of certain characteristics of data type hierarchies, then we have to change entities from one data type into another. 
The general syntax for converting a value val to a type T is T(val).
*/

/*
var geek1 int = 845

// explicit type conversion
var geek2 float64 = float64(geek1)

var geek3 int64 = int64(geek1)

var geek4 uint = uint(geek1)

*/

package main
 
import "fmt"
 
func main() {
 
    // taking the required
    // data into variables
    var totalsum int = 846
    var number int = 19
    var avg float32
 
    // explicit type conversion
    avg = float32(totalsum) / float32(number)
 
    // Displaying the result
	fmt.Printf("Average = %f\n", avg)
	
/*
var geek1 int64 = 875

// it will give compile time error as we
are performing an assignment between mixed types i.e. int64 as int type

var geek2 int = geek1

var geek3 int = 100

// it gives compile time error
as this is invalid operation because types are mix i.e. int64 and int

var addition = geek1 + geek3
fmt.Println(addition)
*/
}