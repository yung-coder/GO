package main

import "fmt"

func main() {
	days := []string{"Sunday", "Tuesday", "Wednesday", "Friday"}

	fmt.Println(days)

	for d := 0; d < len(days); d++ {
		fmt.Println(days[d])
	}

	for i := range days {
		fmt.Println(days[i])
	}

	// most used syntax 
	
	for index, day := range days {
		fmt.Printf("index is %v and val is %v\n", index, day)
	}

	rougeValue := 1

	for rougeValue < 10 {
		if rougeValue == 2 {
			goto lco
		}
		if rougeValue == 5 {
			rougeValue++
			continue
		}
		fmt.Println("Value is:", rougeValue)
		rougeValue++
	}

lco:
	fmt.Println("Jump")
}
