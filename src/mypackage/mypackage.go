package mypackage

import "fmt"

// Car con acceso publico
type CarPublic struct {
	Brand string
	Year  int
}

type CarPrivate struct {
	brand string
	year  int
}

func PrintMessage(s string) {
	fmt.Println("Hola ", s)
}

func printMessage(s string) {
	fmt.Println("Hola ", s)
}
