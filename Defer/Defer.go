package main

import "fmt"

func main() {
	// throws in last line lifo order
	defer fmt.Println("Defer")
	defer fmt.Println("stack")
	defer fmt.Println("queue")

	fmt.Println("Is it working")

	myDefer()
}

func myDefer() {

	// stack type
	for i := 0; i < 5; i++ {
		defer fmt.Print(i)
	}
}
