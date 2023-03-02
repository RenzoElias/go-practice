package main

import (
	"fmt"
	"math"
	"reflect"
)

type figure2D interface {
	getArea() float64
}

type square struct {
	base float64
}

type rectangle struct {
	high  float64
	width float64
}

type trapezoid struct {
	baseA float64
	baseB float64
	high  float64
}

type circle struct {
	radio float64
}

func (s square) getArea() float64 {
	return s.base * s.base
}

func (r rectangle) getArea() float64 {
	return r.high * r.width
}

func (t trapezoid) getArea() float64 {
	return ((t.baseA + t.baseB) / 2) * t.high
}

func (c circle) getArea() float64 {
	return math.Pi * math.Pow(c.radio, 2)
}

func calculateArea(f figure2D) {
	fmt.Printf("Area of %s: %.2f\n", reflect.TypeOf(f).Name(), f.getArea())
}

// CAMBIAR POR MAIN
// func main() {
func main25() {
	mySquare := square{base: 5}
	myRectangle := rectangle{width: 4, high: 5}
	myTrapezoid := trapezoid{baseA: 18, baseB: 10, high: 5}
	myCircle := circle{radio: 4}
	calculateArea(mySquare)
	calculateArea(myRectangle)
	calculateArea(myTrapezoid)
	calculateArea(myCircle)
}

// EXPLAIN
// Tomo a las interfaces más como herramientas para hacer parámetros de Structs en funciones universales, me explico:

// Creamos esta función para calcular el area de cuadrado

// func (c cuadrado) area() float64{
// 	return c.base * c.base
// }
// Normalmente tendríamos que invocar el método del struct cada que queramos usar la función

// myCuadrado := cuadrado{base:2}
// myCuadrado.area()
// Pero gracias a las interfaces, ahora podemos pasar el Struct myCuadrado como parametro y reutilizar una misma funcion en multiples Structs

// //Creamos un Struct triangulo
// type triangulo struct{
// 	base float64
// 	altura float64
// }

// //Instanciamos el struct triangulo
// myTriangulo := triangulo{base: 2, altura: 4}

// //Creamos la interface
// type figuraGeometrica2d interface{
// 	area() float64
// }

// //Ahora tenemos una función que toma como parametro Structs
// func calcularArea(f figuraGeometrica2d){
// 	fmt.Println("El area de tu figura es de: ", f.area)
// }

// // Y ahora, gracias a la interface, podemos pasar a todo Struct que tenga un método area() como parametro de la funcion calcularArea

// calcularArea(myCuadrado)
// calcularArea(myTriangulo)
