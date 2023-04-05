package main

import (
	"fmt"
)

func main() {

	languages := make(map[string]string)

	languages["JS"] = "JavaScript"
	languages["GO"] = "GO"
	languages["TS"] = "TypeScript"
	languages["Dart"] = "Dart"

	fmt.Println("List of all languages", languages)
	fmt.Println("JS  shorst for", languages["JS"])

	delete(languages, "TS")
	fmt.Println(languages)

	// loops in golang

	for key, value := range languages {
		fmt.Printf("For Key %v, value is %v \n", key, value);
	}
}
