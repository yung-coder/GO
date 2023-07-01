package main

import (
	"encoding/json"
	"fmt"
	"github.com/jcelliott/lumber"
	"os"
	"sync"
)

const version = "1.0.0"

type (
	Logger interface {
		Fatal(string, ...interface{})
		Error(string, ...interface{})
		Warn(string, ...interface{})
		Info(string, interface{})
		Debug(string, ...interface{})
		Trace(string, ...interface{})
	}

	Driver struct {
		mutex   sync.Mutex
		mutexes map[string]*sync.Mutex
		dir     string
		log     Logger
	}
)

type Options struct {
   Logger
}


func New()() {
   
}


func (d* Driver) write() error {

}

func (d* Driver) Read () error {

}

func (d* Driver) ReadAll()() {

}

func (d* Driver) Delet() error {

}


func (d* Driver) getorCreateMitex() *sync.Mutex {

}

type Address struct {
	City    string
	State   string
	Country string
	Pincode json.Number
}

type User struct {
	Name    string
	Age     json.Number
	Contact string
	Company string
	Address Address
}

func main() {
	dir := "./"

	db, err := New(dir, nil)
	if err != nil {
		fmt.Println("Error", err)
	}

	employees := []User{
		{"Jhon", "23", "696969699", "Infosys", Address{"palampur", "hp", "india", "14352"}},
		{"Jogi", "26", "6964569699", "Infosysh", Address{"delhi", "it", "india", "14332"}},
		{"rock", "28", "696964799", "Infosysbb", Address{"up", "gp", "india", "14344"}},
	}

	for _, value := range employees {
		db.write("users", value.Name, User{
			Name:    value.Name,
			Age:     value.Age,
			Contact: value.Contact,
			Company: value.Company,
			Address: value.Address,
		})
	}

	records, err := db.ReadAll("users")
	if err != nil {
		fmt.Println("Error", err)
	}
	fmt.Println(records)

	allusers := []User{}

	for _, f := range records {
		employeeFound := User{}
		if err := json.Unmarshal([]byte(f), &employeeFound); err != nil {
			fmt.Println("Errors", err)
		}
		allusers = append(allusers, employeeFound)
	}

	if err := db.Delet("user", "jhon"); err != nil {
		fmt.Println("Error", err)
	}

	if err := db.DeletAll("user", ""); err != nil {
		fmt.Println("Error", err)
	}

}
