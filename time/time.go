package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Time in go")

	presentime := time.Now()
	fmt.Println(presentime)

	// readable time
	fmt.Println(presentime.Format("01-02-2006"))

	createDate := time.Date(2020, time.February, 10, 23, 23, 0, 0, time.Local)
	fmt.Println(createDate)
	fmt.Println(createDate.Format("01-02-2006"))
}
