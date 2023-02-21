package main

import "fmt"
 
func main() {
    // calling the function
    // function returns two values which are
    // assigned to mul and blank identifier
    mul, _ := mul_div(100, 5)
 
    // only using the mul variable
    fmt.Println("Resultado de 105 x 7 = ", mul)
}

// function returning two 
// values of integer type
func mul_div(n1 int, n2 int) (int, int) {
    // returning the values
    return n1 * n2, n1 / n2
}