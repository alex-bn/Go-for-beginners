package main

import "fmt"

func main() {
	// 1)
	// the length of an array cannot change
	var ages [3]int = [3]int{20, 25, 30}
	// go infers the type automatically for us
	var shortHand = [3]int{22, 33, 44}
	fmt.Println(shortHand)
	
	names := [4]string{"mario", "peach", "bowser", "yoshi"}
	names[1] = "luigi"
	
	fmt.Println(ages, len(ages))
	fmt.Println(names, len(names))


	// 2)
	// slices (use arrays under the hood) and the length can change
	var scores = []int{100, 15, 60}
	scores[2] = 24

	// append will return a new slice, so if we want to update it with the new slice we have to say 
	scores = append(scores, 76)

	fmt.Println(scores, len(scores))

	// slice ranges
	// it's inclusive of the first number but not of the second
	rangeOne := names[1:3]
	// from position 2 'till the end
	rangeTwo := names[2:]
	// first 3 items
	rangeThree := names[:3]

	fmt.Println(rangeOne)
 	fmt.Println(rangeTwo)
 	fmt.Println(rangeThree)

	 rangeOne = append(rangeOne, "koopa")
	 fmt.Println(rangeOne)
}