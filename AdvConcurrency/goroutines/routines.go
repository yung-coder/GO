package main

import (
	"fmt"
	"time"
)

func someFunc(num string) {
	fmt.Println(num)
}

func main() {
	// async operation 
	go someFunc("1")

	time.Sleep(time.Second * 2)
	fmt.Println("Routine")
}
