package main

import "fmt"

func message(text string, c chan string) {
	// text se asignara al canal
	c <- text
}

func main29() {
	c := make(chan string, 2)
	// llenar los 2 canales // limitados en el make como tope 2
	c <- "Mensaje 1"
	c <- "Mensaje 2"

	// len - cuantas gorutines (datos) hay dentro ocupados en ese channel
	// cap - capacidad o cantidad maxima que puede almacenar ese channel // en este caso su tope es 2 por el make
	fmt.Println(len(c), cap(c))

	// Range y close
	close(c) // cerrar el canal - y a pesar que tuviera mas espacio para agregar ya no se podra porque esta cerrado
	//c<-"Mensaje 3" // Dara error porque esta cerrado y tmb por que esta limitado a 2 entradas por el make

	// Recorrido de los mensajes que estan dentro de ese channel (2 datos o goroutines dentro del channel)
	for message := range c {
		fmt.Println(message)
	}

	// Select
	email1 := make(chan string)
	email2 := make(chan string)

	go message("--mensaje 1", email1)
	go message("--mensaje 2", email2)

	for i := 0; i < 2; i++ { // es hasta 2 porque solo hay dos canales
		select {
		// Esto es para saber cual canal responde primero (salida)
		// Para saber cual responde primero(salida) se usa select
		case m1 := <-email1:
			fmt.Println("Email recibido de email 1", m1)
		case m2 := <-email2:
			fmt.Println("Email recibido de email 2", m2)
		}
	}
}
