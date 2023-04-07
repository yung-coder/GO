package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"coursename"`
	Price    int
	Platform string   `json:"website"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("JSON handling")
	// EncodeJson()
	DecodeJson()
}

func EncodeJson() {
	lcoCourses := []course{
		{"Reactjs Bootcamp", 299, "Learnit", "abc456", []string{"web-dev", "js"}},
		{"MERN Bootcamp", 199, "Learnit", "abc457", []string{"mern", "js"}},
		{"Flutter Bootcamp", 299, "Learnit", "abc459", nil},
	}

	// package this data as JSON data

	finalJson, err := json.MarshalIndent(lcoCourses, "", "\t")

	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", finalJson)

}

func DecodeJson() {
	jsonWebdata := []byte(`
	{
		"coursename": "MERN Bootcamp",
		"Price": 199,
		"website": "Learnit",
		"tags": [
				"mern",
				"js"
		]
	}
	`)

	var lco course

	chekvalid := json.Valid(jsonWebdata)

	if chekvalid {
		fmt.Println("Json is valid")
		json.Unmarshal(jsonWebdata, &lco)
		// printing struct
		fmt.Printf("%#v\n", lco)
	} else {
		fmt.Println("Not valid")
	}

	// in from of key value pair

	var mymapdata map[string]interface{}

	json.Unmarshal(jsonWebdata, &mymapdata)
	fmt.Printf("%#v\n", mymapdata)

	for k, v := range mymapdata {
		fmt.Printf("Key is %v and value is %v and Type is: %T\n", k, v, v)
	}
}
