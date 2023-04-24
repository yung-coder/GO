package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

var MySigningKey = []byte(os.Getenv("SECRET_KEY"))

func GetJWT() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["client"] = "chad"
	claims["aud"] = "billing.go"
	claims["iss"] = "jwtgo"
	claims["exp"] = time.Now().Add(time.Minute * 1).Unix()

	tokenString, err := token.SignedString(MySigningKey)

	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func Index(w http.ResponseWriter, r *http.Response) {
	validToken, err := GetJWT()
	fmt.Println(validToken)
	if err != nil {
		fmt.Println("Failed to generate token")
	}
	fmt.Println(w, string(validToken))
}

func handelRequest() {
	http.HandleFunc("/", Index)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	handelRequest()
}
