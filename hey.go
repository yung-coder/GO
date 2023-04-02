package main

import "fmt"

func main() {
	var user string = "sak"
	fmt.Println(user)
	fmt.Printf("Var is of type: %T \n", user)

	var isDone bool = true
	fmt.Println(isDone)
	fmt.Printf("Var is of type: %T \n", isDone)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Var is of type: %T \n", smallVal)

	var smallfloatVal float64 = 255.4545454123123
	fmt.Println(smallfloatVal)
	fmt.Printf("Var is of type: %T \n", smallfloatVal)

	// implicit type 

	var randomvar = "Jazz";
	fmt.Println(randomvar);

	// no var style  -- go style 

	dammm := 4499999;
	fmt.Println(dammm);
}
