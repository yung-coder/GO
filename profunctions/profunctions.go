package main

import "fmt"

func main() {

	proResult, msg := proAdder(2, 5, 8, 7)
	fmt.Println("Pro result", proResult)
	println(msg)
}

func proAdder(values ...int) (int, string) {
	total := 0

	for _, val := range values {
		total += val
	}

	return total, "It's a pro function"
}
