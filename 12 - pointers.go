package main

import "fmt"

// func updateName(x string) {
// x = "wedge"
func updateName(x *string) {
	*x = "wedge"
}

func main() {
	name := "mario"

	// updateName(name)

	// fmt.Println("memory address of name is: ", &name)
	
	m := &name
	fmt.Println("memory address:", m)
	fmt.Println("value at memory address:", *m)
	// fmt.Println(name)

	fmt.Println(name)
	updateName(m)
	fmt.Println(name)

}

/*
|--name---|----m----|
|  0x001  |  0x002  |
|---------|---------|
| "mario" | p0x001  |
|---------|---------|
*/