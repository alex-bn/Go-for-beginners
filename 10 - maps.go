package main

import "fmt"

func main() {

	// maps allow us to store key-value pairs where the keys can be different types and the values can be different types as well
	menu := map[string]float64{
		"soup":           4.99,
		"pie":            7.99,
		"salad":          6.99,
		"toffee pudding": 3.55,
	}

	fmt.Println(menu)
	fmt.Println(menu["pie"])

	// looping maps
	for k, v := range menu {
		fmt.Println(k, "-", v)
	}

	// ints as key type
	phoneBook := map [int]string {
		122354345: "mario",
		122354445: "luigi",
		122354323: "bowser",
	}

	fmt.Println(phoneBook)
	fmt.Println(phoneBook[122354345])

	// update the map with another value
	phoneBook[122354323] = "peach"
	fmt.Println(phoneBook)

}
