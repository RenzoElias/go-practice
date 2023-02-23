package main

import (
	"fmt"
	"math"
)

func main() {

	n := 2 == 3
	m := 5 == 5

	fmt.Printf("%v, %T\n", n, n)
	fmt.Printf("%v, %T\n", m, m)

	a := 5
	b := 3
	fmt.Println("===", a)
	fmt.Println(a & b)  // and
	fmt.Println(a | b)  // or
	fmt.Println(a ^ b)  // xor - alt + 94
	fmt.Println(a &^ b) // and or - alt + 94
	fmt.Println("===", b)

	y := 5
	fmt.Println(y << 4) // se añade 4 al exponente
	fmt.Println(y >> 2) // se añade 2 al exponente

	// =========================================
	// =========================================

	fmt.Println()
	fmt.Println("Tipos de datos primitivos")
	fmt.Println("*************************")
	fmt.Println()

	fmt.Println("Especificar el tipo hace que el rendimiento sea mayor a dejarlo declarado implícitamente")
	fmt.Println("La declaración de los tipos es en minúsculas (No existe String por ejemplo al estilo del lenguaje Java")
	fmt.Println()

	fmt.Println("Textos")
	fmt.Println("------")
	fmt.Println()

	fmt.Println("string = \"mi textico\" >>> el texto se coloca entre comillas dobles")
	fmt.Println()

	fmt.Println("Booleano - lógicos")
	fmt.Println("------------------")
	fmt.Println()

	fmt.Println("bool = true o false >>> Solo esas dos opciones")
	fmt.Println()

	fmt.Println("Números enteros")
	fmt.Println("---------------")
	fmt.Println()

	fmt.Println("int8 >>> -2^7 a 2^7-1 >>>", math.MinInt8, "a", math.MaxInt8)
	fmt.Println("int16 >>> -2^15 a 2^15-1 >>>", math.MinInt16, "a", math.MaxInt16)
	fmt.Println("int32 >>> -2^31 a 2^31-1 >>>", math.MinInt32, "a", math.MaxInt32)
	fmt.Println("int64 >>> -2^63 a 2^63-1 >>>", math.MinInt64, "a", math.MaxInt64)
	fmt.Println()

	fmt.Println("Números enteros positivos")
	fmt.Println("-------------------------")
	fmt.Println()

	fmt.Println("uint8 >>> 0 a 2^8-1 >>>", math.MaxUint8)
	fmt.Println("uint16 >>> 0 a 2^16-1 >>", math.MaxUint16)
	fmt.Println("uint32 >>> 0 a 2^32-1 >>>", math.MaxUint32)
	fmt.Println("uint64 >>> 0 a 2^64-1 >>>", uint64(math.MaxUint64))
	fmt.Println()

	fmt.Println("Números enteros positivos")
	fmt.Println("-------------------------")
	fmt.Println()

	fmt.Println("float32 >>> +/- 1.18e^-38 a +/- 3.4e^38 >>>", math.SmallestNonzeroFloat32, "a", math.MaxFloat32)
	fmt.Println("float64 >>> +/- 2.23e^-308 a +/- 1.8e^308 >>>", math.SmallestNonzeroFloat64, "a", math.MaxFloat64)
	fmt.Println()

	fmt.Println("Números complejos")
	fmt.Println("-------------------------")
	fmt.Println()

	fmt.Println("Complex64 = Real e imaginario, ambos en float32")
	fmt.Println("Complex128 = Real e imaginario, ambos en float64")
	fmt.Println()
	fmt.Println("Se declara: c := 10 + 8i")

	variable := 10 + 8i
	fmt.Println("Su valor es:", variable)

	fmt.Println()

}
