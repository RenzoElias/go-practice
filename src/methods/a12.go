package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {

	//login
	user := "Pedro"
	password := "12345"

	if user == "Pedro" && password == "12345" {
		fmt.Println("Logged in")
	} else if user == "Pedro" {
		fmt.Println("Password incorrect")
	} else if password == "12345" {
		fmt.Println("User incorrect")
	} else {
		fmt.Println("Credencials incorrect")
	}

	// =======================
	if isPair(6) {
		fmt.Println("Number is pair")
	} else {
		fmt.Println("Number is odd")
	}
	if isValidUser("Alpha5", "MyPassword") {
		fmt.Println("Credentials are valid")
	} else {
		fmt.Println("Credentials aren't valid")
	}

	// =========================================
	// =========================================
	//Converting strings to integers
	value, err := strconv.Atoi("23") // Atoi is equivalent to ParseInt(s, 10, 0), converted to type int.

	if err != nil {
		log.Fatal(err) //Fatal is equivalent to Print() followed by a call to os.Exit(1).
	}
	fmt.Println(value)
}

func isPair(num int) bool {
	return num%2 == 0
}

func isValidUser(userName, pass string) bool {
	return userName == "Alpha" && pass == "MyPassword"
}
