package main

import "fmt"

func main() {
	x := 0

	// similar to a while loop
	for x < 5 {
		fmt.Println("value of x is:", x)
		x++
	}

	// traditional for loop
	for i := 0; i < 5; i++ {
		fmt.Println("value of i is:", i)
	}

	// loop through a slice
	names := []string{"mario", "luigi", "yoshi", "peach"}

	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}

	// range loop, a bit like for in
	for index, value := range names {
		fmt.Printf("the value at index %v is %v \n", index, value)
	}

	// when we only want the value and not the index
	for _, value := range names {
		fmt.Printf("the value is %v \n",value)
	}

	fmt.Println(names)
	
}