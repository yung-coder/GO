package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://lco.dev/learn?coursename=reactjs&paymentid=dfddshggh"

func main() {
	fmt.Println(myurl)

	// parsing
	result, _ := url.Parse(myurl)

	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)

	qparmas := result.Query()
	fmt.Printf("The typw of query params are: %T\n", qparmas)

	fmt.Println(qparmas["coursename"])

	for _, val := range qparmas {
		fmt.Println("Params is:", val)
	}

	// to construct a url

	partsofurl := &url.URL{
		Scheme:  "https",
		Host:    "lco.dev",
		Path:    "/tutcss",
		RawPath: "user=hitesh",
	}

	anotherurl := partsofurl.String()
	fmt.Println(anotherurl)
}
