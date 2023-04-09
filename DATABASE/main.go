package main

import (
	"db/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("API FOR MONGO")
	r := router.Router()
	fmt.Println("Server is started")
	log.Fatal(http.ListenAndServe(":3000", r))
	fmt.Println("Listening at port 3000 ....")
}
