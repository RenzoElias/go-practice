package main

import "fmt"

func main() {

   // Ejemplo basico de funcion anonima
   var num int = 0

   incrementar := func() int {
      num++
      return num
   }

   fmt.Println("Num: ", num)  // 0
   fmt.Println("Num: ", incrementar())  // 1
   fmt.Println("Num: ", incrementar())  // 2
}

// Una función anónima es una función definida internamente dentro de un bloque de código, y que no 
// tiene identificador o nombre. Este tipo de funciones no son reutilizables como paquetes, siendo 
// utilizadas únicamente dentro del bloque de código en el que son declaradas.
// El siguiente código declara una función anónima dentro de la función main(), la cual sólo tiene visibilidad dentro de esta función.
