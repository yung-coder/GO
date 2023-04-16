package main

import (
	"fmt"
)

func main() {
	mychanndel := make(chan string)

	go func() {
		mychanndel <- "data"
	}()

	msg := <-mychanndel

	fmt.Println(msg)

}
