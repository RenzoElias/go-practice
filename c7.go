package main

import "fmt"

func main() {
	// Defer: deja para ultimo momento, ejecuta antes de terminar el programa
	// se usa por ejemplo, para cerrar una conexion con una BD, un archivo
	// o para cerrar una conexion de red
	defer fmt.Println("Hola")
	fmt.Println("Mundo")

	// Continue y break
	for i := 0; i < 10; i++ {

		// Continue
		if i == 2 {
			fmt.Println("es 2")
			continue // el Println de abajo no se ejecuta, vuelve al for
		}

		fmt.Println(i)

		// Break
		if i == 5 {
			break
		}
	}
}