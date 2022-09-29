package main

import "fmt"
 
func main() {
	
    n := 2 == 3
    m := 5 == 5

	fmt.Printf("%v, %T\n", n, n)
	fmt.Printf("%v, %T\n", m, m)

	a := 5
	b := 3
	fmt.Println( a & b ) // and
	fmt.Println( a | b ) // or
	fmt.Println( a ^ b ) // xor - alt + 94
	fmt.Println( a &^ b ) // and or - alt + 94

	y := 5
	fmt.Println( y << 4 ) // se añade 4 al exponente
	fmt.Println( y >> 2 ) // se añade 2 al exponente


}

