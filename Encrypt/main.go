package main

import (
	"bytes"
	"fmt"
	"os"

	"golang.org/x/term"

	"encrypt/filecrypt"
)

func main() {
	if len(os.Args) < 2 {
		printHelp()
		os.Exit(0)
	}

	function := os.Args[1]

	switch function {
	case "help":
		printHelp()
	case "encrypt":
		encryptHandle()
	case "decrypt":
		decryptHandle()
	default:
		fmt.Println("Run it")
		os.Exit(1)
	}
}

func printHelp() {
	fmt.Println("Encrypt your files using this program")
}

func encryptHandle() {
	if len(os.Args) < 3 {
		println("missing the path to the file")
		os.Exit(0)
	}
	file := os.Args[2]
	if !validateFile(file) {
		panic("File not Found")
	}

	password := getPassword()
	fmt.Println("\nEncrypting...")
	filecrypt.Encrypt(file, password)
	fmt.Println("\n file successfully protected")
}

func decryptHandle() {

	if len(os.Args) < 3 {
		println("missing the path to the file")
		os.Exit(0)
	}
	file := os.Args[2]
	if !validateFile(file) {
		panic("File not Found")
	}
	fmt.Print("Enter password")
	password, _ := term.ReadPassword(0)
	fmt.Println("\nDecrypting...")
	filecrypt.Decrypt(file, password)
	fmt.Println("\n file successfully protected")

}

func getPassword() []byte {
	fmt.Print("Enter password")
	password, _ := term.ReadPassword(0)
	fmt.Print("\n  Confirm Password: ")
	password2, _ := term.ReadPassword(0)
	if !validatePassword(password, password2) {
		fmt.Print("\n Password do not match, please try again\n")
		return getPassword()
	}
	return password
}

func validatePassword(password1 []byte, password2 []byte) bool {
	if !bytes.Equal(password1, password2) {
		return false
	}
	return true
}

func validateFile(file string) bool {
	if _, err := os.Stat(file); os.IsNotExist(err) {
		return false
	}
	return true
}
