package main

import (
	"fmt"

	pk "github.com/RenzoElias/go-practice/src/mypackage"
)

type pc struct {
	ram   int
	disk  int
	brand string
}

// ======================>=============
// PUNTEROS
func (namePC pc) ping() {
	fmt.Println(namePC.brand, "Pong")
}

// *pc - Vamos acceder a los valores del struct mediante el puntero
func (otherNamePC *pc) duplicateRAM() {
	otherNamePC.ram = otherNamePC.ram * 2
}

// ======================>=============
// STRINGERS
func (mPC pc) String() string {
	return fmt.Sprintf("Tengo %d GB Ram, %d GB Disco y es una %s", mPC.ram, mPC.disk, mPC.brand)
}

func main() {

	var myCar pk.CarPublic
	myCar.Brand = "Ferrari"
	myCar.Year = 2020
	fmt.Println(myCar)

	var myOtherCar pk.CarPrivate
	fmt.Println(myOtherCar)

	pk.PrintMessage("Mundo")

	// Privado por la primera palabra Minuscula
	// Sea en Struct y tmb (tipos de datos dentro del struct)
	// Funciones, Slice, Maps
	// pk.printMessage("Mundo")

	// ===================================
	// PUNTEROS
	a := 50
	b := &a
	// &a - acceso a la direccion de memoria (0x14000122020)
	// *b - acceso al valor de la direccion de memoria (50)

	fmt.Println(a, b, *b)

	*b = 100
	fmt.Println(a) // 100

	myPC := pc{ram: 16, disk: 200, brand: "msi"}
	fmt.Println(myPC)

	// la funcion recibira el struct ya instanciado osea myPC
	myPC.ping()

	// myPC.ram = 32 // Mejor usar la funcion duplicateRAM
	fmt.Println(myPC)
	myPC.duplicateRAM()

	fmt.Println(myPC)
	myPC.duplicateRAM()

	fmt.Println(myPC)

	// ===================================
	// STRINGERS
	myPC1 := pc{ram: 16, brand: "msi", disk: 100}
	// PARA EL STRUCT pc SERA PERSONALIZADO SU PRINT - POR EL METODO
	fmt.Println(myPC1)

}
