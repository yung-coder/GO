package main

import "fmt"

func main() {
	fmt.Println("Pointers")

	var ptr *int

	fmt.Println("Pointer is ", ptr)

	myNumber := 23

	var newptr = &myNumber
	fmt.Println("Value  od actual pointer ", newptr)
	fmt.Println("Value  od actual pointer ", *newptr)

	*newptr = *newptr * 2
	fmt.Println("Val", myNumber)
}
