package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	// "math/rand"
)

func main() {
	// math random
	// rand.Seed(time.Now().UnixNano())
	// fmt.Println(rand.Intn(5) + 1);

	// crypto math
	myRandomNum, _ := rand.Int(rand.Reader, big.NewInt(5))
	fmt.Println(myRandomNum)

}
