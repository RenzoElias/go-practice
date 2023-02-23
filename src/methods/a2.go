package main

import (
	"fmt"
	"reflect"
)

func main() {

	// := // se usa para declarar una variable nueva
	// = // para reasignar el valor de una variable
	cd := "cadena de texto"
	i := 8
	f := 4.56789

	fmt.Println("Esto es una línea de texto", cd)
	fmt.Printf("Esto es una %s\n", cd)
	fmt.Printf("Hay %d manzanas\n", i)
	fmt.Printf("%f\n", f)

	// Declaracion de constantes
	const pi float64 = 3.14
	const pi2 = 3.16

	fmt.Println("Pi:", pi)
	fmt.Println("Pi2:", pi2)

	//Declaracion de varaibles
	base := 12          //Primera forma
	var altura int = 14 //Segunda forma
	var area int        //Go no compila si las variables no son usadas

	// SOLUCIONES
	// var _ = area // Asignar un identificador en blanco
	fmt.Println(base, altura, area) // Printearlo

	//Zero values
	//Go asigna vaalores a variables vacías
	// osea no hay necesidad de declarar una variable con Null o None
	var a int
	var b float64
	var c string
	var d bool

	fmt.Println(a, b, c, d)

	//Ejercicio
	//Calcular el áera del cuadrado
	const baseCuadrado = 10
	areaCuadrado := baseCuadrado * baseCuadrado

	fmt.Println("El área del cuadrado es:", areaCuadrado)

	age := 23
	fmt.Println(reflect.TypeOf(age))

}
