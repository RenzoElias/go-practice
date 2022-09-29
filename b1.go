package main

import "fmt"

func main() {

	// Square area
	const squareBase = 10
	squareArea := squareBase * squareBase

	fmt.Println("Square area: ", squareArea)

	x:= 10
	y:= 50

	// Addition
	result:= x + y
	fmt.Println("Addition: ", result)

	// Subtraction
	result = y - x
	fmt.Println("Subtraction: ", result)

	// Multiplication
	result = x * y
	fmt.Println("Multiplication: ", result)

	// Division
	result = y / x
	fmt.Println("Division: ", result)

	// Modulus
	result = y % x
	fmt.Println("Modulus: ", result)

	// Increment
	x++
	fmt.Println("Increment: ", x)

	// Decrement
	x--
	fmt.Println("Decrement: ", x)
}