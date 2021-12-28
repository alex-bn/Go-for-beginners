package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	greeting := "hello there friends!"

	// strings package
	fmt.Println(strings.Contains(greeting, "hello")) // true
	fmt.Println(strings.Contains(greeting, "hello!")) // false
	// the original value is unchanged
	fmt.Println((strings.ReplaceAll(greeting, "hello", "h1")))
	fmt.Println("original string value =", greeting)
	fmt.Println(strings.ToUpper(greeting))
	fmt.Println(strings.Index(greeting, "ll"))
	fmt.Println(strings.Split(greeting, " "))
	
	// sort package
	ages := []int{45, 67, 10, 50, 76, 60, 50, 25}
	// will alter the original slice
	sort.Ints(ages)
	fmt.Println(ages)
	index := sort.SearchInts(ages,25)
	fmt.Println(index)
	names := []string{"yoshi", "mario", "peach", "bowser", "luigi"}
	sort.Strings(names)
	fmt.Println(names)
	fmt.Println(sort.SearchStrings(names, "bowser"))
}