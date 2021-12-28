// Go is a Pass-by-value Language
// Go makes 'copies' of values when passed into functions
package main

import "fmt"

func updateName(x string) {
	x = "wedge"
}

func updateMenu(y map[string]float64){
	y["coffee"] = 2.99
}

func main() {
// group A types - Non-Pointer Values -> strings, ints, bools, floats, arrays, structs
	name := "tifa"

	updateName(name)

	fmt.Println(name)

// group B types - Pointer Wrapper Values -> slices, maps, functions
menu := map[string]float64{
	"pie": 5.95,
	"ice cream": 3.99,
}

updateMenu(menu)
fmt.Println(menu)



}