package main

import (
	"fmt"
	"time"
)

func main() {
	// var str string = "cadena"
	// var entero int = 3
	// var flotante float64 = 3.15654
	
	// Formas de Uso
	// var str string // global variable // declaring. // default ""
	// var str string = "bugz" // declaring.
	// str := "bugz" // declaring. ":="

	// str = "newbugz" // assigning. "="

	// str := "Una cadena de texto"
	str := "123456"
	arr := []int{10, 20, 30, 40}
	arr = append(arr, 50)
	// Defining t and u for Sub method
    t1 := time.Date(2022, 9, 21, 23, 00, 00, 00, time.UTC)
	t2 := time.Now()
	
	// Calling Sub method
    subtract := t2.Sub(t1)
  
    // Prints output
    fmt.Printf("t-d = %v\n", subtract)

	// len1 := len(str)
	len1 := len(arr)
	
	// using blank identifier 
	// fix err // i declared but not used
	// i:= 8
	// _ = i

	if false {
        str = "bugz"
    }else{
        str = "new bugz"
    }

	fmt.Println(len1)
	fmt.Printf("Valor: %v - Tipo: %T\n", str, str)
}