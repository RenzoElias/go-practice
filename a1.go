package main

import "fmt"

func main() {

	// Nombres válidos
	// Empiezan con una letra o un guion bajo '_'
	// _miVariable1, miVariable2 := doubleReturn(2)
	// fmt.Println("_miVariable1: ", _miVariable1)

	// Nombres no válidos
	// Aquellos que empiezan con carácteres especiales exceptuando el '_' guion bajo o con números

	var i int
	fmt.Println("¿Cuál es tu edad? ")
	fmt.Scanf("%d", &i)
	fmt.Printf("Tienes %d años\n", i)
}