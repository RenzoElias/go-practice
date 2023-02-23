package main

import "fmt"

func main()  {
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
}

func isPair(num int) bool {
	return num % 2 == 0
}

func isValidUser(userName, pass string) bool {
	return userName == "Alpha" && pass == "MyPassword"
}