package main

import "fmt"

type car struct {
	brand string
	year  int
}

type usuario struct {
	nombre, sexo string
	edad, dni    int
}

type Person struct {
	name        string
	last        string
	phoneNumber uint32
}

func (p Person) getName() string {
	return p.name
}

func main() {

	myCar := car{
		brand: "Toyota",
		year:  2018,
	}
	fmt.Println("Datos del auto:", myCar)

	//Otra forma
	var otherCar car
	otherCar.brand = "Toyota"
	otherCar.year = 2018
	fmt.Println(otherCar)

	daniel := Person{
		name:        "Daniel",
		last:        "Franc",
		phoneNumber: 4061022,
	}
	println(daniel.last)

	daniel.last = "Franco"

	println(daniel.last)

	println(daniel.getName())
}
