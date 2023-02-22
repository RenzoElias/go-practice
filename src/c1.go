package main

import "fmt"

//Creando funciones
func calculateRectangleArea(b, h int) {
	rectangleArea := b * h
	fmt.Println("El área del rectángulo es:", rectangleArea)
}

func calculateTrapezeArea(b1, b2, h1 float64) {
	trapezeArea := (b1 + b2) * h1 / 2
	fmt.Println("El área del trapecio es:", trapezeArea)
}

func calculateCircleArea(r float64) {
	const pi float64 = 3.14
	circleArea := pi * (r * r)
	fmt.Println("El área del circulo es:", circleArea)
}

//En main epezará a correr en cónsola
func main() {
	calculateRectangleArea(5, 10)
	calculateTrapezeArea(15, 12, 6)
	calculateCircleArea(10)
}