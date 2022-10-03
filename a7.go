package main

import (
	"fmt"
)

func suma(x int, y int) int {
	var sum int
	sum = x + y
	return sum
}

func sumaDouble(x int, y int) (a int, b int) {
	var sum int
	sum = x + y
	return sum*2, sum*2
}
 
func main() {
	
   fmt.Printf("Resultado de la suma %v\n, y el dato es de tipo %T\n", suma(9,13), suma(9,13))

   // Si no te interesa un valor, se le pone el piso _
   // value, _ := sumaDouble(9,13)
   value, valueAuxForType := sumaDouble(9,13)
   fmt.Printf("\nResultado de la suma %v\n, y el dato es de tipo %T\n\n", value, valueAuxForType)

}

// Constantes
// Const <nombre de la constante> = <valor>
// La declaración de las contantes no está definida con la sintaxis “:=”