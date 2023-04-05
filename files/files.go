package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	content := "This will go in file"

	file, err := os.Create("./files/mygofile.txt")

	if err != nil {
		panic(err)
	}

	length, err := io.WriteString(file, content)

	if err != nil {
		panic(err)
	}
	fmt.Print(length)
	defer file.Close()
	readFile("./files/mygofile.txt")
}

func readFile(filename string) {
	databyte, err := ioutil.ReadFile(filename)

	if err != nil {
		panic(err)
	}

	fmt.Print("Text Data \n", string(databyte))
}
