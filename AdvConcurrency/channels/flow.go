package main

import (
	"fmt"
)

func main() {
	mychanndel := make(chan string)
	mychanndel2 := make(chan string)

	go func() {
		mychanndel <- "data"
	}()

	go func() {
		mychanndel2 <- "data2"
	}()

	select {
	case msgmychanndel := <-mychanndel:
		fmt.Println(msgmychanndel)
	case msgmychanndel2 := <-mychanndel2:
		fmt.Println(msgmychanndel2)
	}

}
