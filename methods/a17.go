package main

import (
	"fmt"
	"strings"
)

func isPalindrome(text string) bool {
	var reverseText string
	var lowerText string = strings.ToLower(text)
	fmt.Println(lowerText)
	for i := len(lowerText) - 1; i >= 0; i-- {
		reverseText += string(lowerText[i])
	}
	if reverseText == lowerText {
		return true
	} else {
		return false
	}
}

func main() {
	var slice = []string{"Hola", "que", "tal"}
	//For range
	for i, value := range slice {
		fmt.Printf("%d %s\n", i, value)
	}

	if isPalindrome("Amor a roma") {
		fmt.Println("Es palíndromo")
	} else {
		fmt.Println("No es palíndromo")
	}

	countryCapitalMap := map[string]string{"France": "Paris", "Italy": "Rome", "Japan": "Tokyo"}

	/* print map using keys*/
	for country := range countryCapitalMap {
		fmt.Println("Capital of", country, "is", countryCapitalMap[country])
	}

}
