package main

import (
	"fmt"
	"sort"
)

func main() {
	var fruitlist = []string{"Apple", "Tomato", "Peach"}
	fmt.Println("Simple list ", fruitlist)

	// adding in the list
	fruitlist = append(fruitlist, "Mango", "Banana")
	fmt.Println(fruitlist)

	fmt.Println("Colon Syntax")
	// colon syntax
	fruitlist = append(fruitlist[1:3]) // --> between 0 and 3
	fmt.Println(fruitlist)

	sort.Strings(fruitlist)
	fmt.Println(fruitlist)

	// how to remove a value from slices based on index

	var courses = []string{"reactjs", "javascript", "swift", "rust"}

	fmt.Println(courses)
	var index int =2;
	courses = append(courses[:index] , courses[index+1:]...);
	fmt.Println(courses);
}
