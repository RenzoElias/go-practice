package main

import (
	"fmt"
	"reflect"
)

type object struct {
	property1 int
	property2 string
}
 
func main() {
	
   z:= make ([]int, 10)
   v:= new(object)

   fmt.Printf("\n\nLa estructura de datos almacenada en z -> %v\n", z)
   fmt.Printf("El tipo de dato -> %v\n\n", reflect.TypeOf(z))

   fmt.Printf("El puntero almacenado en v -> %v\n", v)
   fmt.Printf("Lo que almacena la direccion en v -> %v\n", *v)
   fmt.Printf("El tipo de dato -> %v\n\n", reflect.TypeOf(v))


}

// •Make(): Función que únicamente crea y retorna slices, mapas y canales que serán
// almacenados dentro de una variable.Esta función no retorna información inicializada
// con ceros, mas, sin embargo, go automáticamente inicializa las nuevas declaraciones
// de variables con cero

// •New(): Función que asigna nueva memoria a todo tipo de datos. Recibe un tipo de
// dato y retorna la dirección del tipo de dato creado. Esta función, en comparación con la
// anterior si regresa la información inicializada con cero.
