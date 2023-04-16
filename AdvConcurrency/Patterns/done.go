package main

import (
	"fmt"
	"time"
)

func work(done <-chan bool) {
	for {
		select {
		case <-done:
			return
		default:
			fmt.Println("ON")

		}
	}
}

func main() {
	done := make(chan bool)

	go work(done)

	time.Sleep(time.Second * 3)

	close(done)
}
