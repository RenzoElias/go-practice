package main

// FUNCIONES CON PARAMETROS
import (
	"fmt"
	"math"
)

func suma(x int, y int) int {
	var sum int
	sum = x + y
	return sum
}

func sumaDouble(x int, y int) (a, b int) {
	var sum int
	sum = x + y
	return sum * 2, sum * 2
}

/* Funcion simple sin parametros ni return */
func funcionSimple() {
	fmt.Println("Este es el uso de una funcion simple")
}

/* Funcion con parametros declarados con su tipo en cada uno */
func funcionConParametros(a int, b int, c string) {
	fmt.Println("Esta es funcion con parametros", a, b, c)
}

/* Funcion con parametros declarados con su tipo en uno por cada tipo */
func funcionConParametrosAlt(a, b, f int, c, r string) {
	fmt.Println("Esta es funcion con parametros simplificado", a, b, c)
}

/* Funcion con parametros y un return */
func funcionConParametrosYReturnSimple(a int, b int, c string) int {
	fmt.Println("Esta es funcion con parametros y un return simple", a, b, c)
	return a + b
}

/* Funcion con parametros y doble return  */
func funcionConParametrosYReturnDoble(a int, b int, c string) (return1, return2 int) {
	fmt.Println("Esta es funcion con parametros", a, b, c)
	fmt.Println("Se realiza la suma de a y b")
	return a, a + b
}

func areaDeTriangulo(base, altura float64) float64 {
	area := base * altura / 2
	return area
}

func areaDeCirculo(radio float64) float64 {
	area := math.Pow(radio, 2) * math.Pi
	return area
}

func areaDeTrapecio(baseMenor, baseMayor, altura float64) float64 {
	area := (baseMayor + baseMenor) / 2 * altura
	return area
}

// =========================================
// =========================================

func main() {

	fmt.Printf("Resultado de la suma %v\n, y el dato es de tipo %T\n", suma(9, 13), suma(9, 13))

	// Si solo se quiere un valor, de una funcion
	value1, _ := sumaDouble(10, 5)
	fmt.Println("only value1", value1)

	// Si no te interesa un valor, se le pone el piso _
	// value, _ := sumaDouble(9,13)
	value, valueAuxForType := sumaDouble(9, 13)
	fmt.Printf("\nResultado de la suma %v\n, y el dato es de tipo %T\n\n", value, valueAuxForType)

}

// Constantes
// Const <nombre de la constante> = <valor>
// La declaración de las contantes no está definida con la sintaxis “:=”
