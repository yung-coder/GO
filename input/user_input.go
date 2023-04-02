package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter rating  for us")

	// commma ok  || err ok

	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for reading ", input)
	fmt.Printf("Type of %T", input)
}
