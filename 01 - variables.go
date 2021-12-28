// Declare a main package (a package is a way to group functions, and it's made up of all the files in the same directory)
package main

// Standard library package which contains functions for formatting text, including printing to the console
import "fmt"

var someName = ""

// Entry point of our application, there must bu one and only one main() function
// A main function executes by default when you run the main package
func main() {

	fmt.Println("Hello, ninjas!")

	// strings
	var nameOne string = "mario"
	var nameTwo = "luigi"
	var nameThree string

	fmt.Println(nameOne, nameTwo, nameThree)
	
	nameOne = "peach"
	nameThree = "bowser"
	
	fmt.Println(nameOne, nameTwo, nameThree)

	// shorthand version // can be used only inside a function
	nameFour := "yoshi"
	fmt.Println(nameFour)

	// ints 
	var ageOne int = 20
	var ageTwo = 30
	ageThree := 40

	fmt.Println(ageOne, ageTwo, ageThree)

	// bits & memory
	var numOne int8 = 25 // can take values between -128 and 127 
	var numTwo int8 = -128
	var numThree uint8 = 255
	
	fmt.Println(numOne, numTwo, numThree)

	// float
	var scoreOne float32 = -1.5
	var scoreTwo float64 = 4324324324.321312
	scoreThree := 23.99

	fmt.Println(scoreOne, scoreTwo, scoreThree)

}