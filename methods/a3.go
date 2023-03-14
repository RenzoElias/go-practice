package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	// var str string = "cadena"
	// var entero int = 3
	// var flotante float64 = 3.15654

	// Formas de Uso
	// var str string // global variable // declaring. // default ""
	// var str string = "bugz" // declaring.
	// str := "bugz" // declaring. ":="

	// str = "newbugz" // assigning. "="

	// str := "Una cadena de texto"
	str := "123456"
	arr := []int{10, 20, 30, 40}
	arr = append(arr, 50)
	// Defining t and u for Sub method
	t1 := time.Date(2022, 9, 21, 23, 00, 00, 00, time.UTC)
	t2 := time.Now()

	fmt.Println("t1", t1)
	fmt.Println("t2", t2)

	// Calling Sub method
	subtract := t2.Sub(t1)

	// Prints output
	fmt.Printf("t-d = %v\n", subtract)

	// len1 := len(str)
	len1 := len(arr)

	// using blank identifier
	// fix err // i declared but not used
	// i:= 8
	// _ = i

	if false {
		str = "bugz"
	} else {
		str = "new bugz"
	}

	fmt.Println(len1)
	fmt.Printf("Valor: %v - Tipo: %T\n", str, str)

	// =========================================
	// =========================================
	// area del rectangulo
	base := 4
	altura := 2
	println("Area del rectangulo ", base*altura)

	// area del trapecio
	base = 3
	baseAbajo := 2
	areaTrapecio := altura * (base + baseAbajo) / 2
	println("Area del trapecio ", areaTrapecio)

	// area del circulo
	radio := 2
	areaCirculo := math.Pi * math.Pow(float64(radio), 2)
	println(areaCirculo)

	// =========================================
	// =========================================
	x := 10
	y := 50

	// Suma
	result := x + y
	fmt.Println("Suma:", result)

	// Resta
	result = y - x
	fmt.Println("Resta:", result)

	// Multiplicar
	result = x * y
	fmt.Println("Multiplicacion:", result)

	// Dividir
	result = y / x
	fmt.Println("Dividision:", result)

	// Modulo - Residuo
	result = y % x
	fmt.Println("Modulo:", result)

	// Incremental
	x++ //or assign to a var x + 1
	fmt.Println("Incremental:", x)

	// Decremental
	x--
	fmt.Println("Decremental:", x)

}
