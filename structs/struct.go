package main

import "fmt"

func main() {
	fmt.Println("Structs in golang")

	structures := User{"saksham", "godev@gmail.com", true, 16}
	fmt.Println(structures)
	fmt.Printf("ALL the details are: %+v\n", structures)
	fmt.Printf("Name is %v", structures.Name)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
