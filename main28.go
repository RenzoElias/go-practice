package main

import (
	"fmt"
	"net/http"
	"strconv"
)

// Agrega los canales como parametro de entrada en una funcion
// c chan<- string // si se pone asi en el parametro - solo recibira entrada de datos
// c <-chan string // solo sera salida de datos // y para recibirlo en la funcion
// la salida del gorutine o dato del canal, se asignara a text
// text = <-c // aqui se asignaria la salida de dato del canal a la variable text
func get(id int, c chan string) {
	resp, err := http.Get("https://jsonplaceholder.typicode.com/todos/" + strconv.Itoa(id))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// toda esta funcion de fmt, es lo que ingresara(asignara) al canal
	// <- // ENTRADA AL CANAL
	c <- fmt.Sprintf("Status request #%d\n", id)
}

// (mas eficiente, menos que se necesite una optimizacion enorme de codigo entonces se usaria la forma primitiva de gorutines)
// Channel - hace que entre gorutines se puedan comunicar con parametros // forma de organizar las goroutines
func main28() {
	c := make(chan string) // crear un canal - de tipo string // make(chan string, 2) - cantidad limite 2

	for i := 1; i <= 20; i++ {
		// Funcion para asignar datos a un canal
		// este bucle tendra 20 entradas, con sus respectivos datos de entrada
		go get(i, c)
	}

	for i := 1; i <= 20; i++ {
		// <- // SALIDA DEL CANAL
		// Para poder visualizar las salidas // 20 salidas
		fmt.Println(<-c)
	}
}
